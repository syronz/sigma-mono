package response

import "radiusbilling/internal/core/access"

// CheckAccess is a helper for checking the permission for each method
func (r *Response) CheckAccess(resource string) bool {
	return access.Check(r.Engine, r.Context, resource)
}
