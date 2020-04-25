//Entry point api
//Vercion 1 de la implementacion de la api Json

package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

func Apply(route *mux.Router, ctn *registry.Container) {
	//Index
	restClient(route, ctn)
	//REST orders
	orders(route, ctn)

	//Print all the avaibles routes
	printRoutes(route)
}

func orders(route *mux.Router, ctn *registry.Container) {
	//Crear el caso de uso
	orderUseCase := NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))

	list := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		j, err := orderUseCase.ListOrder()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(j)
	}

	create := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		order := &model.Order{}

		_ = json.NewDecoder(r.Body).Decode(order)

		err := orderUseCase.RegisterOrder(order.Number, order.Invoice, order.ClientID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order)
	}
	//Delete a element
	remove := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")

		invoice, ok := params["invoice"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
		}

		err := orderUseCase.DeleteOrder(invoice)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	update := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		order := &model.Order{}

		_ = json.NewDecoder(r.Body).Decode(order)

		err := orderUseCase.UpdateOrder(order.ClientID, order.Number, order.Invoice)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}

	//Routes
	route.HandleFunc("/orders", list).Methods("GET")
	route.HandleFunc("/order", create).Methods("POST")
	route.HandleFunc("/order/{invoice}", remove).Methods("DELETE")
	route.HandleFunc("/order", update).Methods("PUT")
}

func restClient(route *mux.Router, ctn *registry.Container) {
	clienteUseCase := NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))

	list := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		j, err := clienteUseCase.ListClient()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(j)
	}

	create := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		client := &model.Client{}
		_ = json.NewDecoder(r.Body).Decode(client)
		err := clienteUseCase.RegisterClient(client.Name, client.Orders)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(client)
	}

	//TODO: UPDATE ORDERS
	update := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		client := &model.Client{}
		err := clienteUseCase.UpdateClient(client.ID, client.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(client)
	}

	//Delete a element
	remove := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")

		name, ok := params["name"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
		}

		err := clienteUseCase.DeleteClient(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	//Routes
	route.HandleFunc("/clients", list).Methods("GET")
	route.HandleFunc("/client", create).Methods("POST")
	route.HandleFunc("/client/{name}", remove).Methods("DELETE")
	route.HandleFunc("/client", update).Methods("PUT")
}

func printRoutes(appRoutes *mux.Router) {
	appRoutes.Walk(func(route *mux.Route, router *mux.Router, ancestor []*mux.Route) error {
		t, err := route.GetPathTemplate()
		b, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Printf("%-6.6s    %s \n", b, t)
		return nil
	})
}
