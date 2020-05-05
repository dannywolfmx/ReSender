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

func clientRoutes(router *mux.Router, ctn *registry.Container) {
	clienteUseCase := v1.NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))
	s := router.PathPrefix("/client").Subrouter()
	tmpl, err := template.ParseFiles("template/client/list.tmpl", "template/client/orders.tmpl", "template/client/edit.tmpl", "template/client/new.tmpl", "template/layout/footer.tmpl", "template/layout/main.tmpl")

	if err != nil {
		log.Fatal("Fallo carga template client", err)
		return
	}

	list := func(w http.ResponseWriter, r *http.Request) {
		clients, err := clienteUseCase.ListClient()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		tmpl.ExecuteTemplate(w, "list", clients)
	}

	newData := func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Template client/new error")
			return
		}
		tmpl.ExecuteTemplate(w, "new", nil)
	}

	create := func(w http.ResponseWriter, r *http.Request) {
		err := clienteUseCase.RegisterClient(r.FormValue("name"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		http.Redirect(w, r, "/client/list", 302)

	}

	edit := func(w http.ResponseWriter, r *http.Request) {
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
		client := clienteUseCase.GetClient(uint(id))
		tmpl.ExecuteTemplate(w, "edit", client)
	}

	update := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		client := clienteUseCase.GetClient(uint(id))
		client.Name = r.FormValue("name")
		err := clienteUseCase.UpdateClient(client.ID, client.Name, client.Orders)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		http.Redirect(w, r, "/client/list", 302)
	}

	remove := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		clienteUseCase.DeleteClient(uint(id))
		http.Redirect(w, r, "/client/list", 302)
	}

	orders := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		client := clienteUseCase.GetClient(uint(id))
		tmpl.ExecuteTemplate(w, "orders", client)
	}

	s.HandleFunc("/list", list)
	s.HandleFunc("/new", newData)
	s.HandleFunc("/create", create)
	s.HandleFunc("/remove/{id}", remove)
	s.HandleFunc("/edit/{id}", edit)
	s.HandleFunc("/update/{id}", update)
	s.HandleFunc("/{id}/orders", orders)
}
