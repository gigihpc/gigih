package main

import (
	"context"
	"net/http"
)

var keyValue = struct{}{}

func WithParamHandler(p string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query().Get(p)
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), keyValue, v)))
	})
}

func MustAdminHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if v, ok := r.Context().Value(keyValue).(string); ok && v == "admin" {
			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), keyValue, v)))
			println(v)
			return
		}

		http.Error(w, "You must have administrative permissions", http.StatusForbidden)

	})
}

func admin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yes you are admin"))
}

func guest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome guest"))
}

func main() {

	http.Handle("/admin", WithParamHandler("token", MustAdminHandler(http.HandlerFunc(admin))))
	http.Handle("/guest", http.HandlerFunc(guest))

	port := "2017"
	println("open port: " + port)
	http.ListenAndServe(":"+port, nil)
}
