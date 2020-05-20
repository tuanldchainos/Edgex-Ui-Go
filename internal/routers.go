package internal

import (
	"net/http"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/core"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/handler"

	"github.com/gorilla/mux"
)

// InitRestRoutes is router to handler request from client
func InitRestRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(core.LoginUriPath, handler.LoginHandler).Methods(http.MethodGet)
	r.HandleFunc(core.UserLoginPath, handler.UserLoginHandler).Methods(http.MethodPost)
	r.HandleFunc(core.UserLogoutPath, handler.UserLogout).Methods(http.MethodGet)
	r.HandleFunc(core.UserChangePass, handler.UserChangePassHandler).Methods(http.MethodPost)
	r.HandleFunc(core.DevChangePass, handler.DevChangePassHandler).Methods(http.MethodPost)
	r.HandleFunc(core.DevLoginPath, handler.DevLoginHandler).Methods(http.MethodPost)
	r.HandleFunc(core.DevLogoutPath, handler.DevLogout).Methods(http.MethodGet)

	r.HandleFunc(core.DevHomepagePath, handler.DevHomepageHandler).Methods(http.MethodGet)
	r.HandleFunc(core.UserHomepagePath, handler.UserHomepageHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/dev/appservice/list", handler.ListAppServicesProfile).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/dev/config/appservice/{appservice}", handler.PutAppServiceConfig).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/dev/config/coreservice/{coreservice}", handler.PutCoreServiceConfig).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/dev/config/devservice/{devservice}", handler.PutDevServiceConfig).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/user/config/devservice/{devservice}", handler.PutDevServiceConfig).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user/config/appservice/{appservice}", handler.PutAppServiceConfig).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user/restart/service", handler.RestartService).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user/appservice/list", handler.ListAppServicesProfile).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/metadata/device", handler.GetMetadataDevice).Methods(http.MethodGet)

	return r
}
