package main

import (
	internal "Edgex-Ui-Go/internal"
	"Edgex-Ui-Go/internal/configs"
	"Edgex-Ui-Go/internal/core"
	"Edgex-Ui-Go/internal/pkg/usage"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var confFilePath string

	flag.StringVar(&confFilePath, "conf", "", "Specify local configuration file path")

	flag.Usage = usage.HelpCallback
	flag.Parse()

	err := configs.LoadConfig(confFilePath)
	if err != nil {
		log.Printf("Load config failed. Error:%v\n", err)
		return
	}

	r := internal.InitRestRoutes()

	server := &http.Server{
		Handler:      core.GeneralFilter(r),
		Addr:         ":" + strconv.FormatInt(configs.ServerConf.Port, 10),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("EdgeX UI Server Listen On " + server.Addr)

	log.Fatal(server.ListenAndServe())
}
