package handler

import (
	"Edgex-Ui-Go/internal/configs"
	"fmt"

	"github.com/edgexfoundry/go-mod-registry/pkg/types"
	"github.com/edgexfoundry/go-mod-registry/registry"
)

func InitRegistryClientByServiceKey(serviceKey string, needVersionPath bool, Stem string) (registry.Client, error) {
	registryConfig := types.Config{
		Host:       configs.RegistryConf.Host,
		Port:       configs.RegistryConf.Port,
		Type:       configs.RegistryConf.Type,
		ServiceKey: serviceKey,
	}

	if needVersionPath {
		registryConfig.Stem = Stem + configs.RegistryConf.ServiceVersion + "/"
	} else {
		registryConfig.Stem = Stem
	}

	client, err := registry.NewRegistryClient(registryConfig)
	if err != nil {
		return nil, fmt.Errorf("Connection to Registry could not be made: %v", err)
	}
	if !client.IsAlive() {
		return nil, fmt.Errorf("Registry (%s) is not running", registryConfig.Type)
	}
	return client, nil
}

func GetServiceURLviaRegistry(client registry.Client, serviceName string) (string, error) {
	endpoint, err := client.GetServiceEndpoint(serviceName)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s://%s:%v", "http", endpoint.Host, endpoint.Port)
	return url, nil
}
