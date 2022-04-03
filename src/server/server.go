package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ELPanaJose/api-deno-compiler/src/routes"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", routes.IndexRoute)
	r.HandleFunc("/code", routes.GetCode).Methods("GET")
	r.HandleFunc("/code", routes.PostCode).Methods("POST")
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}
	fmt.Printf("Api on port: %s", port)
	http.ListenAndServe(":"+port, r)
}
