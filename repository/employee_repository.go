package repository

import "belajar-golang/model"

type EmployeeRepository interface {
	Insert(employee model.Employee)
	Update(employee model.Employee)
	Delete(employee model.Employee)
	FindAll() (employees []model.Employee)
	FindOne(id string) (employee model.Employee)
}
