package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0/service"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

//ORDERS REST
func Order(r *mux.Router, ctn *registry.Container) {

	//Get a usecase from the container
	containerUsecase := ctn.Resolve("order-usecase").(usecase.OrderUseCase)
	//Crear el caso de uso
	orderServiceApp := service.NewOrderService(containerUsecase)

	//Delete a element
	remove := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")

		idRemove, ok := params["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idRemove)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = orderServiceApp.DeleteOrder(uint(id))
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

		err := orderServiceApp.UpdateOrder(order.ClientID, order.Number, order.Invoice)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}

	//Routes
	r.HandleFunc("/order/{id}", remove).Methods("DELETE")
	r.HandleFunc("/order", update).Methods("PUT")
}
