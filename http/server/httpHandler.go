package app

import (
	"net/http"
	"xks-go/config"
)

// type Secrets struct {
// 	Name string
// }

func (app *App) HandleApp(w http.ResponseWriter, _ *http.Request) {
	// w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.Write([]byte("Hello XKS!"))
}

func (app *App) HandleSecret(w http.ResponseWriter, _ *http.Request) {
	// w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	xksKey := config.AppConfig().Common.Xkskey

	w.Write([]byte(xksKey))
}
