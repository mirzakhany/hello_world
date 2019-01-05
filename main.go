package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var isHealthy = true

func rootHandler(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Hello world , hostname: %s", hostName)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		fmt.Fprint(w, "OK")
	} else {
		http.Error(w, "FAIL", http.StatusInternalServerError)
	}
}

func toggleHealthzHandler(w http.ResponseWriter, r *http.Request) {
	isHealthy = !isHealthy
	fmt.Fprintf(w, "Health status changed to %v", isHealthy)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/health-toggle", toggleHealthzHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
