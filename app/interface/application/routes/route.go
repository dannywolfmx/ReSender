package routes

import (
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func Apply(router *mux.Router, ctn *registry.Container) {
	clientRoutes(router, ctn)
	orderRoutes(router, ctn)
	////	index(router, ctn)
	assets(router, ctn)
	sendMail(router, ctn)
}

func assets(router *mux.Router, ctn *registry.Container) {
	staticDir := "/assets"
	path := router.PathPrefix(staticDir)
	path.Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
}

func index(router *mux.Router, ctn *registry.Container) {
	tmpl := template.Must(template.ParseFiles("template/index/index.tmpl", "template/layout/main.tmpl"))
	router.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		tmpl.ExecuteTemplate(w, "index", nil)
	})
}

func sendMail(router *mux.Router, ctn *registry.Container) {
	router.HandleFunc("/send", func(w http.ResponseWriter, h *http.Request) {
		send("Hola mundo")
	})
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

func send(body string) {
	from := ""
	pass := ""
	to := ""

	msg := "From:" + from + "\n" + "To:" + to + "\n" + "Subject: Hola\n\n" + body

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Println("Smtp error: ", err)
		return
	}
	log.Print("Enviado ")
}
