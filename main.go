package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/arshsuri96/go-services/details"
	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching deatils")
	hostname, err := details.Gethostname() //typical tuple unpacking
	if err != nil {
		panic(err)
	}
	ip := details.GetOutboundIP()
	response := map[string]string{
		"hostname": hostname,
		"ip":       ip.String(), //we added string since the original data type of ip was net.IP and we changed it to string
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server started")

	log.Fatal(http.ListenAndServe(":80", r))
}
