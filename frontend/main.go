package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"../netinfo"
	"net"
)

var backend_url string

func GetBackendNetinfo() string {
	resp, err := http.Get(backend_url)
	if err != nil {
		return fmt.Sprintf("Error: cannot connect to the backend (%s)", resp.Status)
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Backend response:", resp.Status)
	return string(resp_body)
}

func handler_root(w http.ResponseWriter, r *http.Request) {
	version := "1.0"
	fmt.Fprintf(w, "Hello world v%s\r\n", version)
	fmt.Fprintf(w, "Frontend Net info: ")
	json.NewEncoder(w).Encode(netinfo.GetNetwork())
	fmt.Fprintf(w, "Backend Net info: %s", GetBackendNetinfo())
	fmt.Fprintf(w, "\r\n")
	backend_ip, _ := net.LookupIP("backend")
	fmt.Fprintf(w, "Hostname 'backend' resolves to: %s", backend_ip)
}

func handler_health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'Status': 'OK'}")
}

func main() {
	backend_url = os.Getenv("BACKEND_URL")
	if len(backend_url) == 0 {
		panic("BACKEND_URL environment variable missing")
	}

	fmt.Println("Frontend starting up...")

	http.HandleFunc("/", handler_root)
	http.HandleFunc("/health", handler_health)
	http.ListenAndServe(":8081", nil)
}
