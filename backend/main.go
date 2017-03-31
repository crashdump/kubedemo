package main

import (
	"encoding/json"
	"net/http"
	"../netinfo"
	"fmt"
)

func handler_netinfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(netinfo.GetNetwork())
}

func handler_health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'Status': 'OK'}")
}

func main() {
	fmt.Println("Backend starting up...")

	http.HandleFunc("/netinfo", handler_netinfo)
	http.HandleFunc("/health", handler_health)
	http.ListenAndServe(":8080", nil)
}