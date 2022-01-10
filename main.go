package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var res string

func pResponse(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	res = string(reqBody)

}

func gResponse(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(res)
}

func main() {

	port := os.Getenv("PORT")
	//Router initialization
	r := mux.NewRouter()

	//Router handler
	r.HandleFunc("/api/responses", gResponse).Methods("GET")
	r.HandleFunc("/api/responses", pResponse).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
	//log.Fatal(http.ListenAndServe(":3000", r))
}
