package repository

import "algogrit.com/emp-server/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Create(newEmp entities.Employee) (*entities.Employee, error)
}
