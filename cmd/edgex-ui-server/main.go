package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	internal "githup.com/tuanldchainos/Edgex-Ui-Go/internal"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/core"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/helper"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/memory/developMemory"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/memory/userMemory"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/pkg/usage"
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

	helper.LoadServiceUri()

	r := internal.InitRestRoutes()
	userMemory.SetUserPassword()
	developMemory.SetDevPassword()

	server := &http.Server{
		Handler:      core.GeneralFilter(r),
		Addr:         ":" + strconv.FormatInt(configs.ServerConf.Port, 10),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("EdgeX UI Server Listen On " + server.Addr)

	log.Fatal(server.ListenAndServe())
}
