package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

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

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Template order/list error")
			return
		}

		tmpl.Execute(w, clients)

	}

	newData := func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("template/client/new.tmpl")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Template order/new error")
			return
		}

		tmpl.Execute(w, nil)

	}

	create := func(w http.ResponseWriter, r *http.Request) {

		err := clienteUseCase.RegisterClient(r.FormValue("name"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/order/list", 302)

	}

	edit := func(w http.ResponseWriter, r *http.Request) {
		//GetId
		idClient, ok := mux.Vars(r)["id"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idClient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles("template/client/edit.tmpl")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Template order/edit error")
			return
		}
		client := clienteUseCase.GetClient(int64(id))

		tmpl.Execute(w, client)

	}

	update := func(w http.ResponseWriter, r *http.Request) {
		//GetId
		idClient, ok := mux.Vars(r)["id"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idClient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//Get actual order from db

		client := clienteUseCase.GetClient(int64(id))

		//Set values

		client.Name = r.FormValue("name")

		//Save values

		err = clienteUseCase.UpdateClient(client.ID, client.Name, client.Orders)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/order/list", 302)

	}

	remove := func(w http.ResponseWriter, r *http.Request) {
		idCliente, ok := mux.Vars(r)["id"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idCliente)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		clienteUseCase.DeleteClient(int64(id))

		http.Redirect(w, r, "/order/list", 302)

	}

	s.HandleFunc("/list", list)
	s.HandleFunc("/new", newData)
	s.HandleFunc("/create", create)
	s.HandleFunc("/remove/{id}", remove)
	s.HandleFunc("/edit/{id}", edit)
	s.HandleFunc("/update/{id}", update)
}
