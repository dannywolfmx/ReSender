package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func TestClientList(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/clients", nil)
	if err != nil {
		t.Fatalf("no se pudo ejecutar el request %v", err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestClientCreate(t *testing.T) {
	//Esta estructura define un modelo que debe con estado valido
	clientTest := model.Client{
		Name: "Name test",
		Orders: []model.Order{
			{
				Number:  "test number order",
				Invoice: "test invoice order",
			},
		},
	}
	data, _ := json.Marshal(clientTest)

	req, err := http.NewRequest(http.MethodPost, "/client", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("no se pudo ejecutar el request %v", err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	ctn, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}

	Client(r, ctn)
	r.ServeHTTP(rec, req)
	return rec
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
