package router

import (
	"github.com/go-chi/chi"

	handler "xks-go/http/handler"
	app "xks-go/http/server"
)

func New(a *app.App) *chi.Mux {
	reqLog := a.Logger()
	router := chi.NewRouter()

	// r.MethodFunc("GET", "/", app.HandleIndex)
	router.Method("GET", "/api/v1/hello", handler.NewHandler(a.HandleApp, reqLog))
	router.Method("GET", "/api/v1/healthz", handler.NewHandler(a.Health, reqLog))
	router.Method("GET", "/api/v1/secret", handler.NewHandler(a.HandleSecret, reqLog))

	return router
}
