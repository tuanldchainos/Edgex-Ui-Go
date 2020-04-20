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
	DBConf          Database
	StaticProxyConf StaticProxy
	ProxyMapping    map[string]string
	RegistryConf    RegistryConfig
)

type config struct {
	Server       Service        `toml:"Service"`
	DB           Database       `toml:"Database"`
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

type Scheme struct {
	User    string
	Gateway string
}

type Database struct {
	Host     string
	Name     string
	Port     int64
	Username string
	Password string
	Timeout  int64
	Type     string
	Scheme   Scheme
}

type StaticProxy struct {
	CoreDataHost string
	CoreDataPort string

	CoreMetadataHost string
	CoreMetadataPort string

	CoreCommandHost string
	CoreCommandPort string

	CoreExportHost string
	CoreExportPort string

	RuleEngineHost string
	RuleEnginePort string

	SupportLoggingHost string
	SupportLoggingPort string

	SupportNotificationHost string
	SupportNotificationPort string

	SupportSchedulerHost string
	SupportSchedulerPort string
}

type RegistryConfig struct {
	Host               string
	Port               int
	Type               string
	ConfigRegistryStem string
	ServiceVersion     string
}

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
	DBConf = conf.DB
	StaticProxyConf = conf.StaticProxy
	RegistryConf = conf.RegistryConf
	initStaticProxyMapping()
	return nil
}

func initStaticProxyMapping() {

	ProxyMapping = make(map[string]string, 10)

	ProxyMapping[StaticProxyConf.CoreDataHost] = StaticProxyConf.CoreDataPort
	ProxyMapping[StaticProxyConf.CoreMetadataHost] = StaticProxyConf.CoreMetadataPort
	ProxyMapping[StaticProxyConf.CoreCommandHost] = StaticProxyConf.CoreCommandPort
	ProxyMapping[StaticProxyConf.CoreExportHost] = StaticProxyConf.CoreExportPort

	ProxyMapping[StaticProxyConf.RuleEngineHost] = StaticProxyConf.RuleEnginePort

	ProxyMapping[StaticProxyConf.SupportLoggingHost] = StaticProxyConf.SupportLoggingPort
	ProxyMapping[StaticProxyConf.SupportNotificationHost] = StaticProxyConf.SupportNotificationPort
	ProxyMapping[StaticProxyConf.SupportSchedulerHost] = StaticProxyConf.SupportSchedulerPort
}
