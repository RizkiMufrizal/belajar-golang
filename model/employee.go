package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	Id      uint64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Level   string `json:"level"`
	Salary  int    `json:"salary"`
}
