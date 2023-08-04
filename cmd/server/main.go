package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/entities"
)

var employees = []entities.Employee{
	{1, "Gaurav", "LnD", 10001},
	{2, "Anuj", "Cloud", 10002},
	{3, "Misha", "SRE", 20002},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, employees)
	w.Header().Set("content-type", "application/json")

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

	newEmp.ID = len(employees) + 1
	employees = append(employees, newEmp)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
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
