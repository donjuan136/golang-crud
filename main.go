package main

import (
	"go-crud/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	log.Print("server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
