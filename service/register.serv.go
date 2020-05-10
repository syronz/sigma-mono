package service

import (
	"sigmamono/dto"
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	r "sigmamono/internal/core/access/resource"
	"sigmamono/internal/enum/accountstatus"
	"sigmamono/internal/enum/accounttype"
	"sigmamono/internal/enum/nodestatus"
	"sigmamono/internal/param"
	"sigmamono/model"
	"sigmamono/repo"
	"strings"
	"time"
)

// RegisterServ is a two level struct
type RegisterServ struct {
	Engine *core.Engine
}

// ProvideRegisterService is used inside the wire and it is the first level of domain
func ProvideRegisterService(engine *core.Engine) RegisterServ {
	return RegisterServ{Engine: engine}
}

// Register new user
func (p *RegisterServ) Register(register dto.Register) (result dto.Register, err error) {
	register.Company.Expiration = time.Now().AddDate(0, 1, 0)

	original := p.Engine.DB
	tx := p.Engine.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	p.Engine.DB = tx

	licenseService := ProvideLicenseService(p.Engine)
	var licenses []model.License
	if licenses, err = licenseService.GeneratePublic(consts.FreeVersionID, 1); err != nil {
		p.Engine.CheckError(err, "error in generating public serial for registration")
		return
	}

	register.Company.License = licenses[0].Key

	// create company
	companyServ := ProvideCompanyService(repo.ProvideCompanyRepo(p.Engine))
	var company model.Company
	if company, err = companyServ.Save(register.Company); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	result.Company = company

	// create node
	register.Node.CompanyID = company.ID
	register.Node.Code = 101
	register.Node.MachineID = "basic"
	register.Node.Status = nodestatus.Active
	params := param.Param{
		CompanyID: company.ID,
	}
	nodeServ := ProvideNodeService(repo.ProvideNodeRepo(p.Engine))
	var node model.Node
	if node, err = nodeServ.Save(register.Node, params); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	result.Node = node

	// TODO: fix key for each company
	// create bond
	bondSample := model.Bond{
		CompanyID:   company.ID,
		CompanyName: company.Name,
		NodeCode:    node.Code,
		NodeName:    node.Name,
		Key:         "764f4cb02afcc6d000b5a05d05208b152064dad74ada511bb9accf9ee452ae932a072e7cc1c94df84fd1cfcbf32d5b1201ec021a8f5a5605cf743b537398fcf4bc742aba7d684d3efcbf3e75aed0c7b5c9d4cf568def60ec48644372fa22d03d03607f0090df67f2e6441eefa52023f866cf308c9b39374586d2bf0ed24dba9e1fbd35b4f707a7180864f3a8e404a5b08721023fb4c1a5e194914ac50082677f20ce2677a6d9c1dffbfdd2c1af976c03d928778d1b439bdbc412d6f6bfa4da07f03179441e14c91bb43827c1b63d8e8843054749515912748c3cf4d8175027658950f62fcdc3d19b690ed176966c19cc234e7cfb93c064f77e0b6552f7fd16cd3c0113a39172aad52e8d42888f11002240b065bd90d97814e9ef76555145afc9c44a6d7fea10ca2fa319d4fa92c831f87ea71630daa5099bddc8ce765dbde8be241fb4247a96fa39bfb02d49caac81491207769f9e86b0acdfead84e5fc612a8d71d6488d16b03ac924b1fb0cddf90f4b978305396c5ec91438d75cd081ce76f947eb16e8f415f1ed4720306908d5f26177f3673edc221489a4596cc64e62f23a4bd38f107f92f91984eebdb313b2d09cc0878e3854ff575c50364025f0a659c7892bc9f8bb8805f7bb5ec4fe344d58c8404114bb88b0d80a315d19b47553234f5df148ecf70a2a8cd8f8396d9d970bf34f10185bf96c366e7ca607e9f52ef384a3909a2192984b0666e3a1b20107e5f6b0026d4775d82caf7d8eba796271d924840ce4076719344ec2aa8ddb468678f78655af7a7e7844b763321601066fd6271f5edde5d269993cd42346ad515b30ae93940c8328e4ea6134c758f821994d64036cde2ef9577fd9e305dad5798936504adfd98ce8ed0f8fcd84d2737cb65d06d998423207297c1714ef074061352abcd6b130eb188ff93fc952ecb88e2ab5a22bb5fe4289c16386543d11c3708583abb7b78d83deb91d41835898d9cda1f011cbca6c900bfc9a5df550ac59fd8d79165498fb758e1f58aedbc733245e1f2db40b7fdb6d32a013cfce901cf0c9794cfb247605ebaf3774af90ef9881a6c3fe3a6de4ed23a82d6daa970046597fa1dd1620634f8f9db212864af8e49d322243211269de389cdabf6f22c62147aa746d2c808b0df34ca4b58c5d296aae13faf47363395aee637f417fab1017c3d95a4b4523dcc0c",
		MachineID:   "basic",
		Detail:      "",
	}

	bondServ := ProvideBondService(repo.ProvideBondRepo(p.Engine))
	if _, err = bondServ.Save(bondSample); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	// Create role
	// params := param.Param{
	// 	CompanyID: company.ID,
	// 	NodeCode:  node.Code,
	// }
	params.NodeCode = node.Code

	admin := model.Role{
		// CompanyID: company.ID,
		// NodeCode:  node.Code,
		Name: "Admin",
		Resources: strings.Join([]string{
			r.CompanyRead, r.CompanyWrite, r.CompanyExcel,
			r.UserNames, r.UserWrite, r.UserRead, r.UserReport,
			r.NodeRead, r.NodeWrite, r.NodeExcel,
			r.ActivitySelf, r.ActivityAll,
			r.AccountNames, r.AccountRead, r.AccountWrite, r.AccountExcel,
			r.RoleRead, r.RoleWrite,
		}, ", "),
		Description: "admin has all privileges - do not edit",
	}

	// create user
	roleServe := ProvideRoleService(repo.ProvideRoleRepo(p.Engine))
	var role model.Role
	if role, err = roleServe.Create(admin, params); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	// tx.Commit()
	// p.Engine.DB = original

	// register.User.Account.CompanyID = company.ID
	// register.User.Account.NodeCode = node.Code
	register.User.RoleID = role.ID
	register.User.Account.Code = 110001
	register.User.Account.Status = accountstatus.Active
	register.User.Account.Type = accounttype.Asset

	userServ := ProvideUserService(repo.ProvideUserRepo(p.Engine))
	var user model.User
	if user, err = userServ.CreateRollback(register.User, params); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	tx.Commit()
	p.Engine.DB = original

	result.User = user

	return
}
