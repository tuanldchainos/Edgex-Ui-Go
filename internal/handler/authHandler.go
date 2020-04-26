package handler

import (
	"Edgex-Ui-Go/internal/core"
	"Edgex-Ui-Go/internal/domain"
	"html/template"
	"log"
	"net/http"
)

func DevHomepageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("tmpl").ParseFiles("static/pages/devHomepage.html"))
	if err := tmpl.ExecuteTemplate(w, "devHomepage.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("tmpl").ParseFiles("static/pages/login.html"))
	if err := tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	r.ParseForm()
	user.Name = r.PostFormValue("username")
	user.Password = r.PostFormValue("password")

	if user.Name != core.UserName && user.Password != core.UserPass {
		log.Printf("User: %s login failed ", user.Name)
		http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
	}

	session, _ := core.UserStore.Get(r, core.UserSessionSecretKey)
	session.Values["username"] = user.Name
	session.Save(r, w)

	log.Printf("User: %s login ", user.Name)
	http.Redirect(w, r, core.UserHomepagePath, core.RedirectHttpCode)
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

	session, _ := core.DevStore.Get(r, core.DevSessionSecretKey)
	session.Values["devname"] = dev.Name
	session.Save(r, w)

	log.Printf("User: %s login ", dev.Name)
	http.Redirect(w, r, core.DevHomepagePath, core.RedirectHttpCode)
}

func DevLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := core.DevStore.Get(r, core.DevSessionSecretKey)
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := core.UserStore.Get(r, core.UserSessionSecretKey)
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
}
