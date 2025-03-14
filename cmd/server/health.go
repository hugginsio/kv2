package main

import (
	"encoding/json"
	"log"
	"net/http"

	"git.huggins.io/kv2/api"
)

type HealthServer struct {
	Mux  *http.ServeMux
	Port string
}

func ServeHealthEndpoint() {
	srv := &HealthServer{
		Mux:  http.NewServeMux(),
		Port: ":8080",
	}

	srv.Mux.HandleFunc("/health", srv.getHealth)
	if err := http.ListenAndServe(srv.Port, srv.Mux); err != nil {
		log.Fatalln("failed to serve health endpoint:", err)
	}
}

func (hs *HealthServer) getHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	res := api.HealthResponse{
		Status: "UP",
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
