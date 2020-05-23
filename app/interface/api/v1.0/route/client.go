package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0/service"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

func Client(route *mux.Router, ctn *registry.Container) {
	u := ctn.Resolve("client-usecase").(usecase.ClientUsecase)
	clientServiceApp := service.NewClientService(u)

	list := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		j, err := clientServiceApp.ListClient()
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
		err := json.NewDecoder(r.Body).Decode(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = clientServiceApp.RegisterClient(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(client)
	}

	update := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		client := &model.Client{}
		err := clientServiceApp.UpdateClient(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(client)
	}

	//Delete a element
	remove := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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

		if clientServiceApp.DeleteClient(uint(id)) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	//Routes
	route.HandleFunc("/client", create).Methods("POST")
	route.HandleFunc("/clients", list).Methods("GET")
	route.HandleFunc("/client/{id}", remove).Methods("DELETE")
	route.HandleFunc("/client", update).Methods("PUT")
}
