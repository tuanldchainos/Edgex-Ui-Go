package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/helper"
)

const (
	HttpProtocol           = "http"
	OriginHostReqHeader    = "X-Origin-Host"
	ForwardedHostReqHeader = "X-Forwarded-Host"
)

// func ProxyHandler(w http.ResponseWriter, r *http.Request, path string, prefix string) {

// 	var targetAddr string

// 	switch prefix {
// 	case configs.StaticProxyConf.CoreDataPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.CoreDataHost + ":" + configs.StaticProxyConf.CoreDataPort
// 	case configs.StaticProxyConf.CoreMetadataPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.CoreMetadataHost + ":" + configs.StaticProxyConf.CoreMetadataPort
// 	case configs.StaticProxyConf.CoreCommandPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.CoreCommandHost + ":" + configs.StaticProxyConf.CoreCommandPort
// 	case configs.StaticProxyConf.SupportLoggingPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.SupportLoggingHost + ":" + configs.StaticProxyConf.SupportLoggingPort
// 	case configs.StaticProxyConf.SupportNotificationPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.SupportNotificationHost + ":" + configs.StaticProxyConf.SupportNotificationPort
// 	case configs.StaticProxyConf.SupportSchedulerPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.SupportSchedulerHost + ":" + configs.StaticProxyConf.SupportSchedulerPort
// 	case configs.StaticProxyConf.SystemAgentPath:
// 		targetAddr = HttpProtocol + "://" + configs.StaticProxyConf.SystemAgentHost + ":" + configs.StaticProxyConf.SystemAgentPort
// 	}

// 	origin, _ := url.Parse(targetAddr)

// 	director := func(req *http.Request) {
// 		req.Header.Add(ForwardedHostReqHeader, req.Host)
// 		req.Header.Add(OriginHostReqHeader, origin.Host)
// 		req.URL.Scheme = HttpProtocol
// 		req.URL.Host = origin.Host
// 		req.URL.Path = path
// 	}

// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ServeHTTP(w, r)
// }

func ProxyHandler(w http.ResponseWriter, r *http.Request, path string, prefix string) {

	var targetAddr string

	switch prefix {
	case configs.StaticProxyConf.CoreDataPath:
		targetAddr = helper.CoreDataUri
	case configs.StaticProxyConf.CoreMetadataPath:
		targetAddr = helper.MetaDataUri
	case configs.StaticProxyConf.CoreCommandPath:
		targetAddr = helper.CommandUri
	case configs.StaticProxyConf.SupportLoggingPath:
		targetAddr = helper.LoggingUri
	case configs.StaticProxyConf.SupportNotificationPath:
		targetAddr = helper.NotificationsUri
	case configs.StaticProxyConf.SupportSchedulerPath:
		targetAddr = helper.SchedulerUri
	case configs.StaticProxyConf.SystemAgentPath:
		targetAddr = helper.SystemAgentUri
	}

	origin, _ := url.Parse(targetAddr)

	director := func(req *http.Request) {
		req.Header.Add(ForwardedHostReqHeader, req.Host)
		req.Header.Add(OriginHostReqHeader, origin.Host)
		req.URL.Scheme = HttpProtocol
		req.URL.Host = origin.Host
		req.URL.Path = path
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
