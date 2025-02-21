package data

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//import "example.com/goapp/data"

//var employees []Employee

var db *gorm.DB

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees)
	return employees
}

func UpdateEmployee(employee Employee) *Employee {
	var dbEmployee Employee
	err := db.First(&dbEmployee, employee.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	dbEmployee.Age = employee.Age
	dbEmployee.Namn = employee.Namn
	dbEmployee.City = employee.City
	db.Save(&employee)
	return &employee
}

func CreateNewEmployee(employee Employee) *Employee {
	db.Create(&employee)
	return &employee

}

func GetEmployee(id int) *Employee {
	var employee Employee

	err := db.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &employee
}

func Init() {
	db, _ = gorm.Open(sqlite.Open("employees.sqlite"), &gorm.Config{})
	db.AutoMigrate(&Employee{})

	var antal int64
	db.Model(&Employee{}).Count(&antal)

	if antal == 0 {
		db.Create(&Employee{Id: 1, Age: 51, Namn: "Stefan", City: "Test"})
		db.Create(&Employee{Id: 2, Age: 15, Namn: "Oliver", City: "Test"})
		db.Create(&Employee{Id: 3, Age: 21, Namn: "Josefine", City: "Test"})
		db.Create(&Employee{Id: 4, Age: 28, Namn: "Joe", City: "Test"})
	}

	//employees = append(employees, Employee{Id: 1, Age: 51, Namn: "Stefan", City: "Test"})
	//employees = append(employees, Employee{Id: 2, Age: 15, Namn: "Oliver", City: "Test"})
	//employees = append(employees, Employee{Id: 3, Age: 21, Namn: "Josefine", City: "Test"})
	//employees = append(employees, Employee{Id: 4, Age: 28, Namn: "Joe", City: "Test"})
}
