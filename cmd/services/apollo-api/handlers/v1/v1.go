// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"
	"context"

	"github.com/chagasVinicius/apollo/cmd/services/apollo-api/handlers/v1/categorygrp"
	"github.com/chagasVinicius/apollo/internal/core/category"
	"github.com/chagasVinicius/apollo/internal/core/category/stores/categorydb"
	"github.com/chagasVinicius/apollo/kit/web"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *zap.SugaredLogger
	DB  *bun.DB
}

func Ok(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.Respond(ctx, w, "Ok", http.StatusOK)
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// Register category endpoints.
	bgh := categorygrp.Handlers{
		Category: category.NewCore(categorydb.NewStore(cfg.Log, cfg.DB)),
	}

	app.Handle(http.MethodGet, "", "/health", Ok)
	app.Handle(http.MethodGet, version, "/categories", bgh.Query)
	app.Handle(http.MethodGet, version, "/categories/:id", bgh.QueryByID)
	app.Handle(http.MethodPost, version, "/categories", bgh.Create)
}
