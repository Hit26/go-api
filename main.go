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
	fmt.Print(string(reqBody))

}

func main() {

	port := os.Getenv("PORT")
	//Router initialization
	r := mux.NewRouter()

	//Router handler
	r.HandleFunc("/api/responses", pResponse).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
