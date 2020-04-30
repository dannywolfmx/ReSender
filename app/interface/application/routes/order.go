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

func orderRoutes(router *mux.Router, ctn *registry.Container) {
	orderUseCase := v1.NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))
	s := router.PathPrefix("/orders").Subrouter()

	newTemplate := template.Must(template.ParseFiles("template/order/new.tmpl"))
	editTemplate := template.Must(template.ParseFiles("template/order/edit.tmpl"))

	newData := func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("clientid")
		newTemplate.Execute(w, id)
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
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := orderUseCase.DeleteOrder(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ruta := "/client/" + string(id) + "/orders"
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
		editTemplate.Execute(w, order)
	}
	//Guardar cambios
	update := func(w http.ResponseWriter, r *http.Request) {

	}

	s.HandleFunc("/new", newData)
	s.HandleFunc("/create", create)
	s.HandleFunc("/remove/{id}", remove)
	s.HandleFunc("/edit/{id}", edit)
	s.HandleFunc("/update/{id}", update)
}
