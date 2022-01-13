package main

import (
	"fmt"
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
	//fmt.Print(res)
	fmt.Fprintf(w, res)

}

func gResponse(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(res)
	//fmt.Print(res)
	fmt.Fprintf(w, res)
}

func main() {

	port := os.Getenv("PORT")
	//Router initialization
	r := mux.NewRouter()

	//Router handler
	r.HandleFunc("/api/Presponses", gResponse).Methods("GET")
	r.HandleFunc("/api/Gresponses", pResponse).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
	//log.Fatal(http.ListenAndServe(":3000", r))
}
