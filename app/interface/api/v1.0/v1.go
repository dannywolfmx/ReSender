//Entry point api
//Vercion 1 de la implementacion de la api Json

package v1

import (
	"encoding/json"
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
}

func orders(route *mux.Router, ctn *registry.Container) {
	//Crear el caso de uso
	orderUseCase := NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))

	list := func(w http.ResponseWriter, r *http.Request) {

		j, err := orderUseCase.ListOrder()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(j)
	}

	create := func(w http.ResponseWriter, r *http.Request) {

		order := &model.Order{}
		err := orderUseCase.RegisterOrder(order.Number, order.Invoice, order.ClientID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	//Routes
	route.HandleFunc("/orders", list)
	route.HandleFunc("/order", create)

}

func restClient(route *mux.Router, ctn *registry.Container) {
	clienteUseCase := NewClientService(ctn.Resolve("client-usecase").(usecase.ClientUseCase))

	list := func(w http.ResponseWriter, r *http.Request) {

		j, err := clienteUseCase.ListClient()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(j)
	}

	create := func(w http.ResponseWriter, r *http.Request) {
		client := &model.Client{}
		err := clienteUseCase.RegisterClient(client.Name, client.Orders)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	//Routes
	route.HandleFunc("/clients", create)
	route.HandleFunc("/client", list)
}
