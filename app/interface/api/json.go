package api

import (
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func Apply(route *mux.Router, ctn *registry.Container) {
	v1.Apply(route, ctn)
}
