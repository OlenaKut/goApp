package data

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//import "example.com/goapp/data"

//var employees []Employee

var db *gorm.DB

func openMySql(server, database, username, password string, port int) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, server, port, database)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

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
	return true
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

func Init(file, server, database, username, password string, port int) {
	if len(file) == 0 {
		db = openMySql(server, database, username, password, port)
	} else {
		db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	}

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
