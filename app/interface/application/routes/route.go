package routes

import (
	"net/http"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func Apply(router *mux.Router, ctn *registry.Container) {
	clientRoutes(router, ctn)
	orderRoutes(router, ctn)
}

func getIdParramenter(r *http.Request) (int, bool) {
	idCliente, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, false
	}
	return idStringToInt(idCliente)
}

func idStringToInt(idString string) (int, bool) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, false
	}
	return id, true
}
