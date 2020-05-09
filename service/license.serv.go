package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sigmamono/internal/aes"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"strconv"
	"time"
)

// LicenseServ for injecting auth repo
type LicenseServ struct {
	Engine *core.Engine
}

// ProvideLicenseService for license is used in wire
func ProvideLicenseService(engine *core.Engine) LicenseServ {
	return LicenseServ{Engine: engine}
}

// GeneratePublic create serials for public usage
func (p *LicenseServ) GeneratePublic(versionID types.RowID,
	count int) (licenses []model.License, err error) {

	var version model.Version

	versionService := ProvideVersionService(repo.ProvideVersionRepo(p.Engine))
	if version, err = versionService.FindByID(versionID); err != nil {
		return
	}

	var license model.License
	license.Count = count
	if err = license.Validate(action.Create); err != nil {
		p.Engine.Debug(err)
		return
	}

	today := time.Now()
	version.CreatedAt = &today
	version.UpdatedAt = nil
	version.Description = ""
	// TODO: remove inJSON after finishing GeneratePrivate
	inJSON, err := json.Marshal(version)
	if err != nil {
		return
	}

	var encryptedJSON, decryptedJSON string
	_ = inJSON

	for i := 1; i <= count; i++ {
		encryptedJSON, err = aes.EncryptTwice(version.ID.ToString())
		decryptedJSON, err = aes.DecryptTwice(encryptedJSON)
		serial := "0000" + strconv.Itoa(i)
		serial = serial[len(serial)-4:]
		now := time.Now().Format("060102-1504-")
		license := model.License{
			Name:   version.Name,
			Key:    encryptedJSON,
			Serial: fmt.Sprint(now, serial),
		}
		licenses = append(licenses, license)
	}

	_ = decryptedJSON

	return
}

// Update license parse the license's key and if it eligible return company-key
func (p *LicenseServ) Update(license model.License,
	params param.Param) (companyKeyEncrypted string, err error) {
	var companyKey model.CompanyKey
	var decryptedStr string
	var versionID types.RowID
	var version model.Version

	if decryptedStr, err = aes.DecryptTwice(license.Key); err != nil {
		return
	}

	if versionID, err = types.StrToRowID(decryptedStr); err != nil {
		err = fmt.Errorf(term.License_is_not_valid)
		return
	}

	versionService := ProvideVersionService(repo.ProvideVersionRepo(p.Engine))
	if version, err = versionService.FindByID(versionID); err != nil {
		err = fmt.Errorf(term.License_is_not_valid)
		return
	}

	var company model.Company
	companyServ := ProvideCompanyService(repo.ProvideCompanyRepo(p.Engine))
	if company, err = companyServ.FindByID(params.CompanyID); err != nil {
		return
	}

	newExpiration := time.Now().AddDate(0, version.MonthExpire, 0)

	companyKey.CompanyID = company.ID
	companyKey.CompanyName = company.Name
	companyKey.CompanyLegalName = company.LegalName
	companyKey.ServerAddress = p.Engine.Env.Cloud.HostURL
	companyKey.Features = version.Features
	companyKey.NodeCount = version.NodeCount
	companyKey.LocationCount = version.LocationCount
	companyKey.UserCount = version.UserCount
	companyKey.Expiration = newExpiration

	var companyKeyJSON []byte
	if companyKeyJSON, err = json.Marshal(companyKey); err != nil {
		return
	}

	if companyKeyEncrypted, err = aes.EncryptTwice(string(companyKeyJSON)); err != nil {
		return
	}

	company.Expiration = newExpiration
	company.Key = companyKeyEncrypted

	if company, err = companyServ.Save(company); err != nil {
		p.Engine.CheckError(err, "error in saving the company")
		return
	}

	bondServ := ProvideBondService(repo.ProvideBondRepo(p.Engine))
	var bond model.Bond
	if bond, err = bondServ.FindByCompanyID(company.ID); err != nil {
		err = core.NewErrorWithStatus(term.Error_in_casting, http.StatusInternalServerError)
		p.Engine.CheckError(err, "error in bond inside the license.serve.go")
		return

	}

	bond.Key = companyKeyEncrypted

	if bond, err = bondServ.Save(bond); err != nil {
		p.Engine.CheckError(bond.Error, "error in bond inside the license.serve.go")
		return
	}

	return
}
