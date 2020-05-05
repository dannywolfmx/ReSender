package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

func orderRoutes(router *mux.Router, ctn *registry.Container) {
	orderUseCase := v1.NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))
	s := router.PathPrefix("/orders").Subrouter()

	tmpl := template.Must(template.ParseFiles("template/layout/main.tmpl", "template/order/edit.tmpl", "template/order/new.tmpl"))

	newData := func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("clientid")
		tmpl.ExecuteTemplate(w, "new", id)
	}

	create := func(w http.ResponseWriter, r *http.Request) {

		number := r.FormValue("number")
		invoice := r.FormValue("invoice")
		idClient := r.FormValue("id_client")
		id, ok := idStringToInt(idClient)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := orderUseCase.RegisterOrder(number, invoice, uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		ruta := "/client/" + idClient + "/orders"
		http.Redirect(w, r, ruta, 302)
	}

	remove := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		idClient := r.FormValue("clientid")
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := orderUseCase.DeleteOrder(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ruta := fmt.Sprintf("/client/%s/orders", idClient)
		http.Redirect(w, r, ruta, 302)
	}

	//Formulario de edicion
	edit := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		order := orderUseCase.GetOrder(uint(id))
		tmpl.ExecuteTemplate(w, "edit", order)
	}
	//Guardar cambios
	update := func(w http.ResponseWriter, r *http.Request) {
		id, ok := getIdParramenter(r)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		number := r.FormValue("number")
		invoice := r.FormValue("invoice")
		idClient := r.FormValue("id_client")

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := orderUseCase.UpdateOrder(uint(id), number, invoice)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ruta := fmt.Sprintf("/client/%s/orders", idClient)
		http.Redirect(w, r, ruta, 302)
	}

	s.HandleFunc("/new", newData)
	s.HandleFunc("/create", create)
	s.HandleFunc("/remove/{id}", remove)
	s.HandleFunc("/edit/{id}", edit)
	s.HandleFunc("/update/{id}", update)
}
