package core

import (
	"Edgex-Ui-Go/internal/configs"
	"Edgex-Ui-Go/internal/domain"
	"net/http"
	"strings"
)

var DevToken = make(map[string]domain.Dev)

func GeneralFilter(h http.Handler) http.Handler {
	authFilter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	return AuthFilter(authFilter)
}

func AuthFilter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == LoginUriPath || path == "/api/v1/dev/login" || path == "/api/v1/user/login" {
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

		token := GetMd5String(DevelopName)
		_, isValid := DevToken[token]

		if !(isValid) {
			http.Redirect(w, r, LoginUriPath, RedirectHttpCode)
			return
		}

		h.ServeHTTP(w, r)
	})
}
