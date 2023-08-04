package repository

import "algogrit.com/emp-server/entities"

type inmemRepo struct {
	employees []entities.Employee
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 10001},
		{2, "Anuj", "Cloud", 10002},
		{3, "Misha", "SRE", 20002},
	}

	return &inmemRepo{employees}
}
