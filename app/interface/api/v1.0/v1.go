//Entry point api
//Vercion 1 de la implementacion de la api Json

package v1

import (
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0/route"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func Apply(r *mux.Router, ctn *registry.Container) {
	//REST orders
	route.Order(r, ctn)
	route.Client(r, ctn)
}
