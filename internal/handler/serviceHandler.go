package handler

import (
	"Edgex-Ui-Go/internal/configs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Edgex-Ui-Go/internal/core"

	"github.com/pelletier/go-toml"

	"github.com/gorilla/mux"
)

// ListAppServicesProfile return all app service profile
func ListAppServicesProfile(w http.ResponseWriter, r *http.Request) {
	configuration := make(map[string]interface{})
	client, err := InitRegistryClientByServiceKey(configs.RegistryConf.ServiceVersion, false, core.ConfigAppRegistryStem)
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

// PutServiceConfigViaAgent change service config via agent
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

	client, err := InitRegistryClientByServiceKey(coreservice, true, core.ConfigDevRegistryStem)
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

	client, err := InitRegistryClientByServiceKey(appserviceKey, true, core.ConfigAppRegistryStem)
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
	w.Write([]byte("update app service config successfully"))
}
