package core

import (
	"net/http"
	"strings"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"

	"github.com/gorilla/sessions"
)

var DevStore = sessions.NewCookieStore([]byte(DevSessionSecretKey))
var UserStore = sessions.NewCookieStore([]byte(UserSessionSecretKey))

func GeneralFilter(h http.Handler) http.Handler {
	authFilter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	return AuthFilter(authFilter)
}

func AuthFilter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == LoginUriPath || path == DevLoginPath || path == DevLogoutPath || path == UserLoginPath || path == UserLogoutPath || path == UserChangePass {
			h.ServeHTTP(w, r)
			return
		}

		if strings.HasSuffix(path, HtmlSuffix) ||
			strings.HasSuffix(path, CssSuffix) ||
			strings.HasSuffix(path, JsSuffix) ||
			strings.HasSuffix(path, JsonSuffix) ||
			strings.HasPrefix(path, VendorsPath) ||
			strings.HasPrefix(path, DataPathPrefix) {

			http.FileServer(http.Dir(configs.ServerConf.StaticResourcesPath)).ServeHTTP(w, r)
			return
		}

		if path == RootURIPath {
			http.FileServer(http.Dir(configs.ServerConf.StaticResourcesPath)).ServeHTTP(w, r)
			return
		}

		devSession, _ := DevStore.Get(r, DevSessionSecretKey)
		devname := devSession.Values["devname"]

		userSession, _ := UserStore.Get(r, UserSessionSecretKey)
		username := userSession.Values["username"]

		if devname == nil && username == nil {
			http.Redirect(w, r, LoginUriPath, RedirectHttpCode)
			return
		}

		if devname == nil && strings.HasPrefix(path, devPrefix) {
			http.Redirect(w, r, LoginUriPath, RedirectHttpCode)
			return
		}

		if devname != nil {
			for prefix := range configs.ProxyMapping {
				if strings.HasPrefix(path, prefix) {
					path = strings.TrimPrefix(path, prefix)
					ProxyHandler(w, r, path, prefix)
					return
				}
			}
		}

		h.ServeHTTP(w, r)
	})
}
