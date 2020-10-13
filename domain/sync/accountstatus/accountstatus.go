package accountstatus

import "omega/internal/types"

// Account's status
const (
	Active   types.Enum = "active"
	Inactive types.Enum = "inactive"
)

// List used for validation
var List = []types.Enum{
	Active,
	Inactive,
}

// Join make a string for showing in the api
func Join() string {
	return types.JoinEnum(List)
}
