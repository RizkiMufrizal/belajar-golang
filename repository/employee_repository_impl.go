package repository

import (
	"belajar-golang/model"
	"gorm.io/gorm"
)

func NewEmployeeRepository(database *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{DB: database}
}

type EmployeeRepositoryImpl struct {
	DB *gorm.DB
}

func (e EmployeeRepositoryImpl) Insert(employee model.Employee) {
	e.DB.Create(&employee)
}

func (e EmployeeRepositoryImpl) Update(employee model.Employee) {
	e.DB.Where("id = ?", employee.Id).Updates(&employee)
}

func (e EmployeeRepositoryImpl) Delete(employee model.Employee) {
	e.DB.Unscoped().Where("id = ?", employee.Id).Delete(&employee)
}

func (e EmployeeRepositoryImpl) FindAll() (employees []model.Employee) {
	var employee []model.Employee
	e.DB.Find(&employee)
	return employee
}

func (e EmployeeRepositoryImpl) FindOne(id string) (employee model.Employee) {
	var emp model.Employee
	result := e.DB.Find(&emp, id)
	if result.RowsAffected == 0 {
		return model.Employee{}
	}
	return emp
}
