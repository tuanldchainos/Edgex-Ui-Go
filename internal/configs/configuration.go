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
	SystemAgentName string
	SystemAgentHost string
	SystemAgentPort string
}

type RegistryConfig struct {
	Host               string
	Port               int
	Type               string
	ConfigRegistryStem string
	ServiceVersion     string
}

//
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
	ProxyMapping[StaticProxyConf.SystemAgentName] = StaticProxyConf.SystemAgentPort
}
