package main

import (
	"app/api"
	"app/config"
	"app/storage"
)

func main() {

	stg := storage.NewStorage(config.ConnStr)

	defer stg.Close()

	h := api.NewApi(stg)

	h.Run()
}
