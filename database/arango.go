package database

import (
	"context"
	"crypto/tls"
	"den-arango/utils"

	arango "github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
)

// Arangodb is connection to ArangoDB
var Arangodb arango.Client

// DB ...
var DB arango.Database

// Ctx ...
var Ctx context.Context

// Connect ...
func Connect() (err error) {
	conn, err := arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: []string{utils.EnvStr("ARANGO_HOST")},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		return
	}
	Arangodb, err = arango.NewClient(arango.ClientConfig{
		Connection:     conn,
		Authentication: arango.BasicAuthentication(utils.EnvStr("ARANGO_USER"), utils.EnvStr("ARANGO_PASS")),
	})
	if err != nil {
		return
	}
	Ctx = context.Background()
	DB, err = Arangodb.Database(Ctx, utils.EnvStr("ARANGO_DB"))
	if err != nil {
		return
	}
	return nil
}

// BaseModel ...
type BaseModel struct {
	ID  string `json:"_id"`
	Rev string `json:"_rev"`
	Key string `json:"_key"`
}
