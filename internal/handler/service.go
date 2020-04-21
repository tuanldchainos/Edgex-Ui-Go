package handler

import (
	"Edgex-Ui-Go/internal/configs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edgexfoundry/go-mod-registry/pkg/types"
	"github.com/edgexfoundry/go-mod-registry/registry"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/local"
	"github.com/gorilla/mux"
)

const ServiceConfigurableFileName = "configuration.toml"

func ListAppServicesProfile(w http.ResponseWriter, r *http.Request) {
	configuration := make(map[string]interface{})
	client, err := initRegistryClientByServiceKey(configs.RegistryConf.ServiceVersion, false)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	rawConfiguration, err := client.GetConfiguration(&configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	actual, ok := rawConfiguration.(*map[string]interface{})
	if !ok {
		log.Printf("Configuration from Registry failed type check")
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(*actual)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write([]byte(jsonData))
}

// GetServiceConFig return service config, found by service key
func GetServiceConFig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	serviceKey := vars["service"]
	ctx := r.Context()

	client, err := initRegistryClientByServiceKey(serviceKey, true)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	urlServicePrefix, err := getServiceURLviaRegistry(client, serviceKey)
	if err != nil {
		log.Printf(err.Error())
		log.Println("get url via configuration.toml file")
		//http.Error(w, "Can get service url", http.StatusInternalServerError)
		urlServicePrefix = fmt.Sprintf("%s://%s:%v", "http", serviceKey, configs.ProxyMapping[serviceKey])
	}

	var url string

	url = urlServicePrefix + "/api/v1/config"
	fmt.Println(url)
	body, err := clients.GetRequestWithURL(ctx, url)
	if err != nil {
		log.Printf(err.Error())
	}
	ReponseHTTPrequest(w, body, err)
}

// PutServiceConfig change service config
func PutServiceConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	serviceKey := vars["service"]
	ctx := r.Context()

	client, err := initRegistryClientByServiceKey(clients.SystemManagementAgentServiceKey, true)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	urlServicePrefix, err := getServiceURLviaRegistry(client, clients.SystemManagementAgentServiceKey)
	if err != nil {
		log.Printf(err.Error())
		log.Println("get url via configuration.toml file")
		//http.Error(w, "Can get service url", http.StatusInternalServerError)
		urlServicePrefix = fmt.Sprintf("%s://%s:%v", "http", serviceKey, configs.ProxyMapping[serviceKey])
	}

	urlPre := local.New(urlServicePrefix)
	urlPath := "/api/v1/config/" + serviceKey
	urlBody, _ := json.Marshal(r.Body)

	res, err := clients.PutRequest(ctx, urlPath, urlBody, urlPre)
	if err != nil {
		log.Printf(err.Error())
	}
	ReponseHTTPrequest(w, []byte(res), err)
}

func initRegistryClientByServiceKey(serviceKey string, needVersionPath bool) (registry.Client, error) {
	registryConfig := types.Config{
		Host:       configs.RegistryConf.Host,
		Port:       configs.RegistryConf.Port,
		Type:       configs.RegistryConf.Type,
		ServiceKey: serviceKey,
	}

	if needVersionPath {
		registryConfig.Stem = configs.RegistryConf.ConfigRegistryStem + configs.RegistryConf.ServiceVersion + "/"
	} else {
		registryConfig.Stem = configs.RegistryConf.ConfigRegistryStem
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

func getServiceURLviaRegistry(client registry.Client, serviceName string) (string, error) {
	endpoint, err := client.GetServiceEndpoint(serviceName)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s://%s:%v", "http", endpoint.Host, endpoint.Port)
	return url, nil
}
