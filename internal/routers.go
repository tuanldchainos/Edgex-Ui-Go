package internal

import (
	"Edgex-Ui-Go/internal/core"
	"Edgex-Ui-Go/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRestRoutes is router to handler request from client
func InitRestRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc(core.LoginUriPath, handler.LoginHandler).Methods(http.MethodGet)
	r.HandleFunc(core.UserLoginPath, handler.UserLoginHandler).Methods(http.MethodPost)
	r.HandleFunc(core.UserLogoutPath, handler.UserLogout).Methods(http.MethodGet)
	r.HandleFunc(core.DevLoginPath, handler.DevLoginHandler).Methods(http.MethodPost)
	r.HandleFunc(core.DevLogoutPath, handler.DevLogout).Methods(http.MethodGet)

	r.HandleFunc(core.DevHomepagePath, handler.DevHomepageHandler).Methods(http.MethodGet)
	//r.HandleFunc(core.UserHomepagePath, handler.HandlerConfig).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/appservice/list", handler.ListAppServicesProfile).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.GetServiceConFig).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.PutServiceConfig).Methods(http.MethodPost)

	return r
}
