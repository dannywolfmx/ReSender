package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dannywolfmx/ReSender/app/interface/application"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

//path sqlite

func main() {
	//Inicializar la base de datos

	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	route := mux.NewRouter()

	//	api.Apply(route, ctn)

	application.Run(route, ctn)
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      route,
	}
	log.Fatal(server.ListenAndServe())

}
