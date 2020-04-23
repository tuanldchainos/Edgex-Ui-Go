package internal

import (
	"Edgex-Ui-Go/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRestRoutes is router to handler request from client
func InitRestRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/auth/login", handler.LoginHandler)
	r.HandleFunc("/userlogin", handler.UserLoginHandler)
	r.HandleFunc("/devlogin", handler.DevLoginHandler)
	r.HandleFunc("/config", handler.HandlerConfig)
	r.HandleFunc("/api/v1/appservice/list", handler.ListAppServicesProfile).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.GetServiceConFig).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.PutServiceConfig).Methods(http.MethodPost)
	return r
}
