package routes

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
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
		files := saveFile(r.MultipartForm.File["archivos"])
		err = orderUseCase.RegisterOrder(number, invoice, m, files, uint(id))
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

func saveFile(files []*multipart.FileHeader) []model.File {
	dataFiles := []model.File{}
	for _, file := range files {

		f, err := file.Open()
		defer f.Close()
		if err != nil {
			panic("Error al abrir el archivo")
		}
		//data info
		data := model.File{
			Title: file.Filename,
			Path:  "./assets/upload/" + file.Filename,
		}

		//Crear un archivo vacio con el nombre
		destino, err := os.Create(data.Path)
		defer destino.Close()

		if err != nil {
			panic("Error al crear contenedor para el archivo")
		}

		_, err = io.Copy(destino, f)
		if err != nil {
			panic("Error al copiar los datos al archivo")
		}
		dataFiles = append(dataFiles, data)
	}
	return dataFiles

}
