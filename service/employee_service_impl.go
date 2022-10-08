package service

import (
	"belajar-golang/model"
	"belajar-golang/repository"
)

func NewEmployeeService(employeeRepository *repository.EmployeeRepository) EmployeeService {
	return &EmployeeRepositoryImpl{EmployeeRepository: *employeeRepository}
}

type EmployeeRepositoryImpl struct {
	EmployeeRepository repository.EmployeeRepository
}

func (e EmployeeRepositoryImpl) Insert(employee model.Employee) {
	e.EmployeeRepository.Insert(employee)
}

func (e EmployeeRepositoryImpl) Update(employee model.Employee) {
	e.EmployeeRepository.Update(employee)
}

func (e EmployeeRepositoryImpl) Delete(employee model.Employee) {
	e.EmployeeRepository.Delete(employee)
}

func (e EmployeeRepositoryImpl) FindAll() (employees []model.Employee) {
	return e.EmployeeRepository.FindAll()
}

func (e EmployeeRepositoryImpl) FindOne(id string) (employee model.Employee) {
	return e.EmployeeRepository.FindOne(id)
}
