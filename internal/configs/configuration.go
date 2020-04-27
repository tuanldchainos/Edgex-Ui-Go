package configs

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	defaultConfigFilePath = "res/configuration.toml"
)

var (
	ServerConf      Service
	StaticProxyConf StaticProxy
	ProxyMapping    map[string]string
	RegistryConf    RegistryConfig
)

type config struct {
	Server       Service        `toml:"Service"`
	StaticProxy  StaticProxy    `toml:"StaticProxy"`
	RegistryConf RegistryConfig `toml:"Registry"`
}

type Service struct {
	Host                string
	Port                int64
	Labels              []string
	OpenMsg             string
	StaticResourcesPath string
}

type StaticProxy struct {
	CoreDataPath string
	CoreDataPort string
	CoreDataHost string

	CoreMetadataPath string
	CoreMetadataPort string
	CoreMetadataHost string

	CoreCommandPath string
	CoreCommandPort string
	CoreCommandHost string

	SupportLoggingPath string
	SupportLoggingPort string
	SupportLoggingHost string

	SupportNotificationPath string
	SupportNotificationPort string
	SupportNotificationHost string

	SupportSchedulerPath string
	SupportSchedulerPort string
	SupportSchedulerHost string

	SystemAgentPath string
	SystemAgentPort string
	SystemAgentHost string
}

type RegistryConfig struct {
	Host           string
	Port           int
	Type           string
	ServiceVersion string
}

// LoadConfig get config form configuration.toml
func LoadConfig(confFilePath string) error {
	if len(confFilePath) == 0 {
		confFilePath = defaultConfigFilePath
	}

	absPath, err := filepath.Abs(confFilePath)
	if err != nil {
		log.Printf("Could not create absolute path to load configuration: %s; %v", absPath, err.Error())
		return err
	}
	log.Printf("Loading configuration from: %s\n", absPath)
	var conf config
	if _, err := toml.DecodeFile(absPath, &conf); err != nil {
		log.Printf("Decode Config File Error:%v", err)
		return err
	}
	ServerConf = conf.Server
	StaticProxyConf = conf.StaticProxy
	RegistryConf = conf.RegistryConf
	initStaticProxyMapping()
	return nil
}

func initStaticProxyMapping() {

	ProxyMapping = make(map[string]string, 10)

	ProxyMapping[StaticProxyConf.CoreDataPath] = StaticProxyConf.CoreDataPort
	ProxyMapping[StaticProxyConf.CoreMetadataPath] = StaticProxyConf.CoreMetadataPort
	ProxyMapping[StaticProxyConf.CoreCommandPath] = StaticProxyConf.CoreCommandPort
	ProxyMapping[StaticProxyConf.SystemAgentPath] = StaticProxyConf.SystemAgentPort
	ProxyMapping[StaticProxyConf.SupportLoggingPath] = StaticProxyConf.SupportLoggingPort
	ProxyMapping[StaticProxyConf.SupportNotificationPath] = StaticProxyConf.SupportNotificationPort
	ProxyMapping[StaticProxyConf.SupportSchedulerPath] = StaticProxyConf.SupportSchedulerPort
}
