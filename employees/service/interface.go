package service

import "algogrit.com/emp-server/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmp entities.Employee) (*entities.Employee, error)
}
