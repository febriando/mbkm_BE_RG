package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(405)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
		}
	}) // TODO: replace this
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(405)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
		}
	}) // TODO: replace this
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(405)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
		}
	}) // TODO: replace this
}
