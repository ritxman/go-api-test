package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get all environment variables
	port := os.Getenv("PORT")
	testSecret := os.Getenv("TEST_SECRET")

	r.HandleFunc("/api/test/{value}", test)
	fmt.Println("Server is running on :", port)
	fmt.Println("Test Secret:", testSecret)
	http.ListenAndServe(":"+port, r)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	vars := mux.Vars(r)
	value := vars["value"]

	response := map[string]string{"hello": value}
	json.NewEncoder(w).Encode(response)
}
