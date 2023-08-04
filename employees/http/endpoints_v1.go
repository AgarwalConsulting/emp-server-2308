package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"

	"algogrit.com/emp-server/entities"
)

func (h EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")

	employees, err := h.v1Svc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func (h EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	validate := validator.New()
	errs := validate.Struct(newEmp)

	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errs)
		return
	}

	createdEmp, err := h.v1Svc.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}
