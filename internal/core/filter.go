package core

import (
	"Edgex-Ui-Go/internal/configs"
	"net/http"
	"strings"

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
		if path == LoginUriPath || path == DevLoginPath || path == DevLogoutPath || path == UserLoginPath || path == UserLogoutPath {
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

		if devname == nil {
			http.Redirect(w, r, LoginUriPath, RedirectHttpCode)
			return
		}

		for prefix := range configs.ProxyMapping {
			if strings.HasPrefix(path, prefix) {
				path = strings.TrimPrefix(path, prefix)
				ProxyHandler(w, r, path, prefix)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}
