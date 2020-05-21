package service

import (
	"errors"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/utils/password"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthServ defining auth service
type AuthServ struct {
	Engine *core.Engine
}

// ProvideAuthService for auth is used in wire
func ProvideAuthService(engine *core.Engine) AuthServ {
	return AuthServ{Engine: engine}
}

// Login User
func (p *AuthServ) Login(auth model.Auth) (user model.User, err error) {
	if err = auth.Validate(action.Login); err != nil {
		return
	}

	jwtKey := []byte(p.Engine.Env.Setting.JWTSecretKey)

	// user = connector.New().
	// 	Domain(domains.Administration).
	// 	Entity("User").
	// 	Method("FindByUsername").
	// 	Args(auth.Username).
	// 	SendReceive(p.Engine).(model.User)

	// if err = user.Error; err != nil {
	// 	err = errors.New(term.Username_or_password_is_wrong)
	// 	return
	// }

	userServ := ProvideUserService(repo.ProvideUserRepo(p.Engine))
	if user, err = userServ.FindByUsername(auth.Username); err != nil {
		err = errors.New(term.Username_or_password_is_wrong)
		return
	}

	if password.Verify(auth.Password, user.Password,
		p.Engine.Env.Setting.PasswordSalt) {

		// station := connector.New().
		// 	Domain(domains.Central).
		// 	Entity("Station").
		// 	Method("FindByCompanyID").
		// 	Args(user.Account.CompanyID).
		// 	SendReceive(p.Engine).(model.Station)

		// companyKey := station.Extra["company_key"].(model.CompanyKey)
		// if companyKey.Expiration.Before(time.Now()) {
		// 	err = errors.New(term.Company_license_has_been_expired)
		// 	return
		// }

		stationServ := ProvideStationService(repo.ProvideStationRepo(p.Engine))
		var station model.Station
		if station, err = stationServ.FindByCompanyID(user.Account.CompanyID); err != nil {
			err = errors.New(term.Company_not_exist_in_station)
			return
		}

		companyKey := station.Extra["company_key"].(model.CompanyKey)
		if companyKey.Expiration.Before(time.Now()) {
			err = errors.New(term.Company_license_has_been_expired)
			return
		}

		expirationTime := time.Now().
			Add(time.Duration(p.Engine.Env.Setting.JWTExpiration) * time.Second)
		claims := &types.JWTClaims{
			Username:  auth.Username,
			ID:        user.ID,
			Language:  user.Language,
			CompanyID: station.CompanyID,
			NodeCode:  station.NodeCode,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var extra struct {
			Token string `json:"token"`
		}
		if extra.Token, err = token.SignedString(jwtKey); err != nil {
			return
		}

		user.Extra = extra
		user.Password = ""

	} else {
		err = errors.New(term.Username_or_password_is_wrong)
	}

	return
}
