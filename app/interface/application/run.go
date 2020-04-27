package application

import (
	"github.com/dannywolfmx/ReSender/app/interface/application/routes"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func Run(server *mux.Router, ctn *registry.Container) {
	routes.Apply(server, ctn)
}
