package handler

import (
	"Edgex-Ui-Go/internal/core"
	"Edgex-Ui-Go/internal/domain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandlerConfig(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("tmpl").ParseFiles("static/pages/config.html"))
	if err := tmpl.ExecuteTemplate(w, "config.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("tmpl").ParseFiles("static/login.html"))
	if err := tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type UserLogin struct {
		username string
		userpass string
	}

	var user UserLogin

	r.ParseForm()
	user.username = r.PostFormValue("username")
	user.userpass = r.PostFormValue("password")

	if user.userpass == "24111998" {
		http.Redirect(w, r, "http://www.google.com", 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func DevLoginHandler(w http.ResponseWriter, r *http.Request) {

	var dev domain.Dev

	r.ParseForm()
	dev.Name = r.PostFormValue("username")
	dev.Password = r.PostFormValue("password")

	if dev.Name != core.DevelopName && dev.Password != core.DevelopPass {
		log.Printf("User: %s login failed ", dev.Name)
		http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
	}

	token := core.GetMd5String(dev.Name)
	core.DevToken[token] = dev

	log.Printf("User: %s login ", dev.Name)
	http.Redirect(w, r, "/api/v1/dev/homepage", core.RedirectHttpCode)
}

func DevLogout(w http.ResponseWriter, r *http.Request) {
	delete(core.DevToken, core.GetMd5String(core.DevelopName))
	fmt.Println(core.DevToken)
	http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
}
