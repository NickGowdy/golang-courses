package main

import (
	"log"
	"net/http"

	"github.com/pluralsight/inventoryservice/product"
)

const basePath = "/api"

func main() {
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
