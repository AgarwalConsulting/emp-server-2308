package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	out := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		// if //
		h.ServeHTTP(w, req)

		dur := time.Since(begin)
		fmt.Printf("%s %s took %s\n", req.Method, req.URL, dur)
	}

	return http.HandlerFunc(out)
}

func main() {
	var empRepo = repository.NewInMem()
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empSvc)

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
}
