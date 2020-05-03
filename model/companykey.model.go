package model

import (
	"sigmamono/internal/types"
	"time"
)

// CompanyKey model used in the activation and after successfully parse license
type CompanyKey struct {
	CompanyID        types.RowID `json:"company_id,omitempty"`
	CompanyName      string      `json:"company_name,omitempty"`
	CompanyLegalName string      `json:"key,omitempty"`
	ServerAddress    string      `json:"server_address,omitempty"`
	Features         string      `json:"features,omitempty"`
	NodeCount        int         `json:"node_count"`
	LocationCount    int         `json:"location_count"`
	UserCount        int         `json:"user_count"`
	Expiration       time.Time   `json:"expiration"`
	Error            error       `json:"error"`
}
