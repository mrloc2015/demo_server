package main

import (
	"fmt"
	"net/http"
)

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	fmt.Printf("We have a request at url: %s\n", message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "./people.json")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	msg := "OK"
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func main() {
	//////////
	fmt.Println("Server listening at 9090")
	http.HandleFunc("/v1/health", healthHandler)
	http.HandleFunc("/v1/people", peopleHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
