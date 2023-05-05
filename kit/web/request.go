package web

import (
	"net/http"
	"fmt"

	"github.com/dimfeld/httptreemux/v5"
)

type validator interface {
	Validate() error
}

// Param returns the web call parameters from the request.
func Param(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

// Query returns the query string value for the provided key.
func Query(r *http.Request, key string, fallback interface{}) string {
	if v := r.URL.Query().Get(key); len(v) > 0 {
		return v
	}
	return fmt.Sprint(fallback)
}

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
// If the provided value is a struct then it is checked for validation tags.
// If the value implements a validate function, it is executed.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
