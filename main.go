package main

import (
	"github.com/google/uuid"
	"net/http"
	"log"
)

var serverUUID = uuid.New().String()

func generateUUIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(serverUUID))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	http.HandleFunc("/uuid", generateUUIDHandler)
	log.Println("Starting server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
