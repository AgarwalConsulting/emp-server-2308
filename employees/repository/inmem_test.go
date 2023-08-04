package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	initialEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(initialEmps))

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			emp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}

			createdEmp, err := sut.Create(emp)

			assert.Nil(t, err)
			assert.NotEqual(t, 0, createdEmp.ID)

			someEmps, err := sut.ListAll()

			assert.Nil(t, err)
			assert.NotNil(t, someEmps)
		}()
	}

	wg.Wait()

	finalEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.Equal(t, 103, len(finalEmps))
}
