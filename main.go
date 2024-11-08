package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/test/{value}", test)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["value"]

	response := map[string]string{"hello": value}
	json.NewEncoder(w).Encode(response)
}
