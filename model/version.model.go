package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
)

// Version model
type Version struct {
	types.GormCol
	Name          string                 `gorm:"not null;unique" json:"name,omitempty"`
	Features      string                 `gorm:"type:text" json:"features,omitempty"`
	NodeCount     int                    `json:"node_count,omitempty"`
	LocationCount int                    `json:"location_count,omitempty"`
	UserCount     int                    `json:"user_count,omitempty"`
	MonthExpire   int                    `json:"month_expire,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Extra         map[string]interface{} `sql:"-" json:"extra_version,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Version) Pattern() string {
	return `(versions.name LIKE '%[1]v%%' OR
		versions.features LIKE '%[1]v%%' OR
		versions.description LIKE '%%%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Version) Columns(variate string) (string, error) {
	full := []string{"versions.id", "versions.name", "versions.features",
		"versions.created_at", "versions.updated_at", "versions.description"}

	return checkColumns(full, variate)
}

// Validate check the type of
func (p *Version) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_versions_form)

	switch act {
	case action.Save:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, term.Name, "name")
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
