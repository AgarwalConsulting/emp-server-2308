package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

var empRepo = repository.NewInMem()
var empSvc = service.NewV1(empRepo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")

	employees, err := empSvc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empSvc.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

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
	// http.DefaultServeMux
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
}
