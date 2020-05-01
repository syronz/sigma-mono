package types

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims for JWT
type JWTClaims struct {
	Username  string `json:"username"`
	ID        RowID  `json:"id"`
	Language  string `json:"language"`
	CompanyID RowID  `json:"company_id"`
	NodeCode  uint64 `json:"node_code"`
	jwt.StandardClaims
}
