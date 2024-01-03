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
	fmt.Println("BU main bo'lmagan BRANCHida yozildi")
	fmt.Println("BU main BRANCHida yozildi")
	h.Run()
}
