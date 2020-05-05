package handler

import (
	"Edgex-Ui-Go/internal/core"
	"Edgex-Ui-Go/internal/domain"
	"Edgex-Ui-Go/internal/userMemory"
	"encoding/json"
	"fmt"
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

func UserHomepageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("tmpl").ParseFiles("static/pages/userHomepage.html"))
	if err := tmpl.ExecuteTemplate(w, "userHomepage.html", nil); err != nil {
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

	var userCre domain.Cre

	r.ParseForm()
	userCre.Name = r.PostFormValue("username")
	userCre.Password = r.PostFormValue("password")
	fmt.Println(userMemory.BasicUser.Pass)
	if userCre.Name != userMemory.BasicUser.Name || userCre.Password != userMemory.BasicUser.Pass {
		log.Printf("User: %s login failed ", userCre.Name)
		http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
	}

	session, _ := core.UserStore.Get(r, core.UserSessionSecretKey)
	session.Values["username"] = userCre.Name
	session.Save(r, w)

	log.Printf("User: %s login ", userCre.Name)
	http.Redirect(w, r, core.UserHomepagePath, core.RedirectHttpCode)
}

func UserChangePassHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userUpdateData map[string]string
	err := json.NewDecoder(r.Body).Decode(&userUpdateData)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	if userUpdateData["oldpass"] == userMemory.BasicUser.Pass {
		userMemory.UpdateUserPass(userUpdateData["newpass"])
		w.Write([]byte("update user password successfully"))
	} else {
		w.Write([]byte("password incorrect, please try again"))
	}
}

func DevLoginHandler(w http.ResponseWriter, r *http.Request) {

	var devCre domain.Cre

	r.ParseForm()
	devCre.Name = r.PostFormValue("username")
	devCre.Password = r.PostFormValue("password")

	if devCre.Name != core.DevelopName || devCre.Password != core.DevelopPass {
		log.Printf("User: %s login failed ", devCre.Name)
		http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
	} else {
		session, _ := core.DevStore.Get(r, core.DevSessionSecretKey)
		session.Values["devname"] = devCre.Name
		session.Save(r, w)

		log.Printf("User: %s login ", devCre.Name)
		http.Redirect(w, r, core.DevHomepagePath, core.RedirectHttpCode)
	}
}

func DevLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := core.DevStore.Get(r, core.DevSessionSecretKey)
	session.Options.MaxAge = -1
	session.Save(r, w)
	// http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := core.UserStore.Get(r, core.UserSessionSecretKey)
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, core.LoginUriPath, core.RedirectHttpCode)
}
