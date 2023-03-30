// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"github.com/chagasVinicius/apollo/kit/web"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *zap.SugaredLogger
	DB  *bun.DB
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	// TODO: add routes
}
