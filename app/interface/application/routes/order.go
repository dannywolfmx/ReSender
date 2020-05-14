package routes

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/interface/api/v1.0"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/dannywolfmx/ReSender/app/usecase"
	"github.com/gorilla/mux"
)

func orderRoutes(router *mux.Router, ctn *registry.Container) {
	orderUseCase := v1.NewOrderService(ctn.Resolve("order-usecase").(usecase.OrderUseCase))

	tmpl := template.Must(template.ParseFiles("template/layout/main.tmpl", "template/order/edit.tmpl", "template/order/new.tmpl"))

	newData := func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("clientid")
		tmpl.ExecuteTemplate(w, "new", id)
	}

	create := func(w http.ResponseWriter, r *http.Request) {

		number := r.FormValue("number")
		invoice := r.FormValue("invoice")
		mails := strings.Split(r.FormValue("mails"), ",")
		idClient := r.FormValue("id_client")
		id, ok := idStringToInt(idClient)

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		m := []model.MailDirection{}
		for _, direccion := range mails {
			m = append(m, model.MailDirection{
				Direction: direccion,
			})
		}

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Save uploaded files
		files := r.MultipartForm.File["archivos"]
		for _, file := range files {

			f, err := file.Open()
			defer f.Close()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//Crear un archivo vacio con el nombre
			destino, err := os.Create("./assets/upload/" + file.Filename)
			defer destino.Close()

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = io.Copy(destino, f)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		err = orderUseCase.RegisterOrder(number, invoice, m, uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
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

	s := router.PathPrefix("/orders").Subrouter()
	s.HandleFunc("/new", newData)
	s.HandleFunc("/create", create)
	s.HandleFunc("/remove/{id}", remove)
	s.HandleFunc("/edit/{id}", edit)
	s.HandleFunc("/update/{id}", update)
}
