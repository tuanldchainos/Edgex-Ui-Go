package core

import (
	"Edgex-Ui-Go/internal/configs"
	"net/http"
	"strings"
)

const (
	RootURIPath  = "/"
	LoginUriPath = "/api/v1/auth/login"
)

const (
	ContentTypeKey   = "Content-Type"
	JsonContentType  = "application/json"
	RedirectHttpCode = 302
	SessionTokenKey  = "X-Session-Token"

	AjaxRequestIdentifier = "XMLHttpRequest"
	AjaxRequestHeader     = "X-Requested-With"

	HtmlSuffix     = ".html"
	CssSuffix      = ".css"
	JsSuffix       = ".js"
	JsonSuffix     = ".json"
	VendorsPath    = "/vendors"
	DataPathPrefix = "/data"

	NoAuthorizationMsg = "no authorization."
)

func GeneralFilter(h http.Handler) http.Handler {
	authFilter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	return AuthFilter(authFilter)
}

func AuthFilter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == LoginUriPath {
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

		if path != LoginUriPath {
			http.Redirect(w, r, LoginUriPath, RedirectHttpCode)
		}

		h.ServeHTTP(w, r)
	})
}
