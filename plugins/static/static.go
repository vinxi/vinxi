package static

import (
	"net/http"

	"gopkg.in/vinxi/vinxi.v0/manager"
)

// New creates a new static plugin who serves
// files of the given server local path.
func New(path string) manager.Plugin {
	return manager.NewPlugin("static", "serve static files", staticHandler(path))
}

func staticHandler(path string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.FileServer(http.Dir(path))
	}
}