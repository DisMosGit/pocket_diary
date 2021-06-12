package user

import (
	"context"

	arango "github.com/arangodb/go-driver"
)

var userTab arango.Collection
var ctx context.Context

// BaseUserModel ...
type BaseUserModel struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

// SaveUserModel ...
type SaveUserModel struct {
	BaseUserModel
}

// GetUserModel ...
type GetUserModel struct {
	arango.DocumentMeta
	BaseUserModel
}
