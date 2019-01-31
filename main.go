package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miles-moran/catStax/shapes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api", solve).Methods("POST")
	fmt.Println("SERVER STARTED AT - PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func solve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Shapes []shapes.PointArray
	_ = json.NewDecoder(r.Body).Decode(&Shapes)
	shapes.Prepare(Shapes)
	json.NewEncoder(w).Encode(Shapes)

}

//TO RUN API
//	go build

// ./catstax

// go build && ./catstax
