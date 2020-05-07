package model

import (
	"fmt"
	"sigmamono/internal/types"
)

// Activity model
type Activity struct {
	types.FixedCol
	Event    string      `gorm:"index:event_idx" json:"event"`
	UserID   types.RowID `json:"user_id"`
	Username string      `gorm:"index:username_idx" json:"username"`
	IP       string      `json:"ip"`
	URI      string      `json:"uri"`
	Before   string      `gorm:"type:text" json:"before"`
	After    string      `gorm:"type:text" json:"after"`
}

// SearchPattern is used for generate query for finding ter  among specific columns
func (p *Activity) SearchPattern(str string) string {
	var pattern = `(activities.event LIKE '%%%[1]v%%' OR
		activities.username LIKE '%[1]v' OR
		activities.id LIKE '%[1]v' OR
		activities.created_at LIKE '%[1]v%%' OR
		activities.data LIKE '%%%[1]v%%' OR
		activities.ip LIKE '%[1]v')`
	return fmt.Sprintf(pattern, str)
}
