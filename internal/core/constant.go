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
	UserChangePass   = "/api/v1/user/change/pass"
)

// server config
const (
	ContentTypeKey   = "Content-Type"
	JsonContentType  = "application/json"
	RedirectHttpCode = 302
	SessionTokenKey  = "X-Session-Token"
	HttpProtocol     = "http"

	AjaxRequestIdentifier = "XMLHttpRequest"
	AjaxRequestHeader     = "X-Requested-With"

	HtmlSuffix     = ".html"
	CssSuffix      = ".css"
	JsSuffix       = ".js"
	JsonSuffix     = ".json"
	DataPathPrefix = "/data"
	VendorsPath    = "/vendors"

	NoAuthorizationMsg     = "no authorization."
	ConfigAppRegistryStem  = "edgex/appservices/"
	ConfigDevRegistryStem  = "edgex/core/"
	ConfigCoreRegistryStem = "edgex/core/"
)

// user and dev login info
const (
	DevelopName = "develop"
	DevelopPass = "develop"
)

// session info
const (
	DevSessionSecretKey  = "DevSession"
	UserSessionSecretKey = "UserSession"

	devPrefix = "/api/v1/dev/"
)

// Edgex service keyword
const (
	ConfigSeedServiceKey            = "edgex-config-seed"
	CoreCommandServiceKey           = "edgex-core-command"
	CoreDataServiceKey              = "edgex-core-data"
	CoreMetaDataServiceKey          = "edgex-core-metadata"
	SupportLoggingServiceKey        = "edgex-support-logging"
	SupportNotificationsServiceKey  = "edgex-support-notifications"
	SystemManagementAgentServiceKey = "edgex-sys-mgmt-agent"
	SupportSchedulerServiceKey      = "edgex-support-scheduler"
)
