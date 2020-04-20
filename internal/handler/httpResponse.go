package handler

import (
	"fmt"
	"net/http"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// ReponseHTTPrequest sending response to client
func ReponseHTTPrequest(w http.ResponseWriter, body []byte, err error) {
	status := getHTTPStatus(err)
	if status != http.StatusOK {
		fmt.Errorf(err.Error())
		http.Error(w, err.Error(), status)
	} else {
		if len(body) > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}
		w.Write(body)
	}
}

func getHTTPStatus(err error) int {
	if err != nil {
		chk, ok := err.(*types.ErrServiceClient)
		if ok {
			return chk.StatusCode
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}