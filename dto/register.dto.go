package dto

import "sigmamono/model"

// Register is a combination of company, node and user
type Register struct {
	Company model.Company `json:"company,omitempty"`
	Node    model.Node    `json:"node,omitempty"`
	User    model.User    `json:"user,omitempty"`
}
