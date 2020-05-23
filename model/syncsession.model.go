package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"time"
)

// SyncSession model
type SyncSession struct {
	types.FixedCol
	CompanyID        types.RowID `gorm:"not null" json:"company_id,omitempty"`
	NodeCode         uint64      `gorm:"not null" json:"node_id,omitempty"`
	MachineID        string      `gorm:"not null" json:"machine_id,omitempty"`
	InitiateAt       *time.Time  `json:"initiate_at,omitempty"`
	NodeToServerAt   *time.Time  `json:"node_to_server_at,omitempty"`
	ServerToNodeAt   *time.Time  `json:"server_to_node_at,omitempty"`
	CloseAt          *time.Time  `json:"close_at,omitempty"`
	Duration         string      `json:"duration,omitempty"`
	NodeToServerSize float64     `json:"node_to_server_size,omitempty"`
	ServerToNodeSize float64     `json:"server_to_node_size,omitempty"`
	NodeToServerRows uint        `json:"node_to_server_rows,omitempty"`
	ServerToNodeRows uint        `json:"server_to_node_rows,omitempty"`
	Status           string      `json:"status,omitempty"`
	Cost             float64     `json:"cost,omitempty"`
	CurrentTime      *time.Time  `sql:"-" json:"current_time,omitempty"`
	LastSyncDate     *time.Time  `sql:"-" json:"last_sync_date,omitempty"`
	Delay            uint        `sql:"-" json:"delay"`
	Error            error       `sql:"-" json:"syncsession_error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p SyncSession) Pattern() string {
	return `(sync_sessions.name LIKE '%[1]v%%' OR
		sync_sessions.id = '%[1]v' OR
		sync_sessions.description LIKE '%%%[1]v%%' OR
		sync_sessions.resources LIKE '%%%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p SyncSession) Columns(variate string) (string, error) {
	full := []string{
		"sync_sessions.id",
		"sync_sessions.company_id",
		"sync_sessions.node_code",
		"sync_sessions.machine_id",
		"sync_sessions.initiate_at",
		"sync_sessions.node_to_server_at",
		"sync_sessions.server_to_node_at",
		"sync_sessions.close_at",
		"sync_sessions.duration",
		"sync_sessions.node_to_server_size",
		"sync_sessions.server_to_node_size",
		"sync_sessions.node_to_server_rows",
		"sync_sessions.server_to_node_rows",
		"sync_sessions.status",
		"sync_sessions.cost",
	}

	return checkColumns(full, variate)
}

// Validate check the type of fields
func (p *SyncSession) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_role_form)

	switch act {
	case action.Save:

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
