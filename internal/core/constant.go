package core

// Root path
const (
	RootURIPath  = "/"
	LoginUriPath = "/api/v1/auth/login"
)

// dev api list
const (
	DevHomepagePath = "/api/v1/dev/homepage"
	DevLoginPath    = "/api/v1/dev/login"
	DevLogoutPath   = "/api/v1/dev/logout"
)

// user api list
const (
	UserHomepagePath = "/api/v1/user/homepage"
	UserLoginPath    = "/api/v1/user/login"
	UserLogoutPath   = "/api/v1/user/logout"
)

// server config
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

// user and dev login info
const (
	DevelopName = "develop"
	DevelopPass = "develop"

	UserName = "user"
	UserPass = "user"
)

// session info
const (
	DevSessionSecretKey  = "DevSession"
	UserSessionSecretKey = "UserSession"
)
