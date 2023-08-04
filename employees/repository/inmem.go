package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmemRepo struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

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

	return &inmemRepo{employees: employees}
}
