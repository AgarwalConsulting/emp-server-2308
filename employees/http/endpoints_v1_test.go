package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockService)

	respRec := httptest.NewRecorder()

	jsonString := `{"name": "Gaurav", "speciality": "LnD", "project": 1}` // Type: string
	reqBody := strings.NewReader(jsonString)                              // io.Reader
	req := httptest.NewRequest("POST", "/v1/employees", reqBody)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1}

	mockService.EXPECT().Create(expectedEmp).Return(&expectedEmp, nil)

	// sut.CreateV1(respRec, req)
	sut.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusOK, respRec.Code)

	var createdEmp entities.Employee
	json.NewDecoder(respRec.Body).Decode(&createdEmp)

	assert.Equal(t, expectedEmp, createdEmp)
}
