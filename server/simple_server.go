package server

import (
	"fmt"
	"html/template"
	"log"
	"nats/storage/cashe"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {

	r := mux.NewRouter()

	r.HandleFunc("/orders", OrderHanlder)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Print(err)
	}

}

func OrderHanlder(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println(id)
	order, ok := cashe.GetOrder(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, order)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
