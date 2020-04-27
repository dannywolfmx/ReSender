package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

func Apply(router *mux.Router, ctn *registry.Container) {
	//server.LoadHTMLGlob("templates/**/*.tmpl")
	clientRoutes(router, ctn)
}

func clientRoutes(router *mux.Router, ctn *registry.Container) {
	clienteUseCase := v1.NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))
	s := router.PathPrefix("/order").Subrouter()

	list := func(w http.ResponseWriter, r *http.Request) {
		clients, err := clienteUseCase.ListClient()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		tmpl, err := template.ParseFiles("template/client/list.tmpl")

		tmpl.Execute(w, clients)

	}

	s.HandleFunc("/list", list)
}
