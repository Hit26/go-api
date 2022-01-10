package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func pResponse(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	p := string(reqBody)
	fmt.Print(p)

}

func main() {

	port := os.Getenv("PORT")
	//Router initialization
	r := mux.NewRouter()

	//Router handler
	r.HandleFunc("/api/responses", pResponse).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
	log.Fatal(http.ListenAndServe(":3000", r))
}
