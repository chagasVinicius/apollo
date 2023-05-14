// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	v1 "github.com/chagasVinicius/apollo/app/services/apollo-api/handlers/v1"
	"github.com/chagasVinicius/apollo/internal/web/v1/mid"
	"github.com/chagasVinicius/apollo/kit/web"
	"github.com/uptrace/bun"
	"github.com/zmb3/spotify/v2"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	DB       *bun.DB
	Tracer   trace.Tracer
	SongClient *spotify.Client
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) http.Handler {
	app := web.NewApp(
		cfg.Shutdown,
		cfg.Tracer,

		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)

	v1.Routes(app, v1.Config{
		Log: cfg.Log,
		DB:  cfg.DB,
		SongClient: cfg.SongClient,
	})

	return app
}
