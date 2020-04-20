package internal

import (
	"Edgex-Ui-Go/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRestRoutes is router to handler request from client
func InitRestRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/config/service/{service}", handler.GetServiceConFig).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.PutServiceConfig).Methods(http.MethodPost)
	return r
}
