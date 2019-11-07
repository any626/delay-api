package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"context"
)

const Timeout = 15 * time.Second

type DelayResponse struct {
	Delay string `json:"delay"`
}

type InvalidDelayResponse struct {
	InvalidDelay string `json:"invalid_delay"`
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), Timeout)
	defer cancel()

	vars := mux.Vars(r)
	delay, err := time.ParseDuration(vars["delay"])
	if err != nil {
		response := InvalidDelayResponse{vars["delay"]}
		invalidResponseInfo(w, response)
		return
	}

	select {
	case <- ctx.Done():
		w.WriteHeader(504)
	case <- time.After(delay):
		response := DelayResponse{fmt.Sprint(delay)}
		responseInfo(w, response)
	}
}

func invalidResponseInfo(w http.ResponseWriter, response InvalidDelayResponse) {
	w.WriteHeader(422)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func responseInfo(w http.ResponseWriter, response DelayResponse) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{delay}", delayHandler)
	http.Handle("/", r)

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}
	log.Fatal(s.ListenAndServe())
}
