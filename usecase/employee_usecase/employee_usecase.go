package employee_usecase

import (
	"EmployeeManagementApp/domain"
	"time"
)

type employeeUsecase struct {
	employeeRepo   domain.EmployeeRepository
	contextTimeout time.Duration
}

func NewEmployeeUsecase(t domain.EmployeeRepository, timeout time.Duration) domain.EmployeeUsecase {
	return &employeeUsecase{
		employeeRepo:   t,
		contextTimeout: timeout,
	}
}
