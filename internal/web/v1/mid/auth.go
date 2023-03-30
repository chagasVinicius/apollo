package mid

import (
	"context"
	"net/http"

	"github.com/chagasVinicius/apollo/internal/web/auth"
	"github.com/chagasVinicius/apollo/kit/web"
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims, err := auth.Authenticate(ctx, r.Header.Get("authorization"))
			if err != nil {
				return auth.NewAuthError("authenticate: failed: %s", err)
			}

			ctx = auth.SetClaims(ctx, claims)

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
