package internal

import (
	"Edgex-Ui-Go/internal/handler"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.New("tmpl").ParseFiles("static/pages/config.html"))

// InitRestRoutes is router to handler request from client
func InitRestRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/config", handlerConfig)
	r.HandleFunc("/api/v1/appservice/list", handler.ListAppServicesProfile).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.GetServiceConFig).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/config/service/{service}", handler.PutServiceConfig).Methods(http.MethodPost)
	return r
}

func handlerConfig(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "You called me!")
	if err := tmpl.ExecuteTemplate(w, "config.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
