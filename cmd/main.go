package main

import (
	"app/api"
	"app/config"
	"app/storage"
	"fmt"
)

func main() {

	stg := storage.NewStorage(config.ConnStr)

	defer stg.Close()

	h := api.NewApi(stg)
	fmt.Println("Bu github sitedan yozildi")
	h.Run()
}
