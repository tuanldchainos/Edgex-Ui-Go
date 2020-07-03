package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/registrySupport"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/core"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/pelletier/go-toml"

	"github.com/gorilla/mux"
)

func ListAppServicesProfile(w http.ResponseWriter, r *http.Request) {
	configuration := make(map[string]interface{})
	client, err := registrySupport.InitRegistryClientByServiceKey(configs.RegistryConf.ServiceVersion, false, core.ConfigAppRegistryStem)
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

func PutCoreServiceConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	coreservice := vars["coreservice"]
	configuration := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	client, err := registrySupport.InitRegistryClientByServiceKey(coreservice, true, core.ConfigCoreRegistryStem)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	configurationTomlTree, err := toml.TreeFromMap(configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	fmt.Println()
	err = client.PutConfigurationToml(configurationTomlTree, true)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("update core service config successfully"))
}

func PutAppServiceConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	appserviceKey := vars["appservice"]
	configuration := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	client, err := registrySupport.InitRegistryClientByServiceKey(appserviceKey, true, core.ConfigAppRegistryStem)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	configurationTomlTree, err := toml.TreeFromMap(configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = client.PutConfigurationToml(configurationTomlTree, true)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("update app service config successfully"))
}

func PutDevServiceConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	devservice := vars["devservice"]
	var configuration map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	client, err := registrySupport.InitRegistryClientByServiceKey(devservice, true, core.ConfigDevRegistryStem)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	fmt.Println(configuration)

	configurationTomlTree, err := toml.TreeFromMap(configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	fmt.Println(configurationTomlTree)
	err = client.PutConfigurationToml(configurationTomlTree, true)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("update device service config successfully"))
}

func RestartService(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.Background()

	agentClient, _ := registrySupport.InitRegistryClientByServiceKey(core.SystemManagementAgentServiceKey, true, core.ConfigCoreRegistryStem)
	agentURI, _ := registrySupport.GetServiceURLviaRegistry(agentClient, core.SystemManagementAgentServiceKey)
	agentURL := agentURI + "/api/v1/operation"
	fmt.Println(agentURL)

	// agentURL := "http://localhost:48090/api/v1/operation"

	configuration := make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&configuration)

	res, _ := clients.PostJSONRequestWithURL(ctx, agentURL, &configuration)
	w.Write([]byte(res))
}

func GetDevServiceConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	devservice := vars["devservice"]
	client, err := registrySupport.InitRegistryClientByServiceKey(devservice, true, core.ConfigDevRegistryStem)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	configuration := make(map[string]interface{})
	configurationData, err := client.GetConfiguration(&configuration)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(configurationData)
	w.Write(data)
}

func GetMetadataDevice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.Background()
	metaClient, _ := registrySupport.InitRegistryClientByServiceKey(core.CoreMetaDataServiceKey, true, core.ConfigCoreRegistryStem)
	URI, _ := registrySupport.GetServiceURLviaRegistry(metaClient, core.CoreMetaDataServiceKey)
	URL := URI + "/api/v1/device"

	res, _ := clients.GetRequestWithURL(ctx, URL)
	w.Write(res)
}
