package handler

import (
	"html/template"
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
	type DevLogin struct {
		devname string
		devpass string
	}

	var dev DevLogin

	r.ParseForm()
	dev.devname = r.PostFormValue("username")
	dev.devpass = r.PostFormValue("password")

	if dev.devpass == "24111998" {
		http.Redirect(w, r, "http://www.google.com", 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
