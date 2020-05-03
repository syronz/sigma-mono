package access

import (
	"sigmamono/internal/core"

	"github.com/gin-gonic/gin"
)

// Check returns true if user hasn't permission otherwise return false
func Check(engine *core.Engine, c *gin.Context, resource string) bool {

	// accessResult := connector.New().
	// 	Domain(domains.Central).
	// 	Entity("Access").
	// 	Method("CheckAccess").
	// 	Args(c, resource).
	// 	SendReceive(engine).(bool)

	return false
}
