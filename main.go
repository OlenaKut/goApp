package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/goapp/data"
	"github.com/gin-gonic/gin"
)

func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)

}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employee := data.GetEmployee(numId)

	if employee == nil { // INTE HITTAT  /api/employee/
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeAdd(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	data.CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}
func apiEmployeeUpdateById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id, _ = strconv.Atoi(id)
	if data.UpdateEmployee(employee) == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

type PageView struct {
	Title  string
	Rubrik string
}

func handleStartPage(c *gin.Context) {
	//c.String(http.StatusOK, "Hello Olena")
	c.HTML(http.StatusOK, "index.html", &PageView{Title: "Hello Olena", Rubrik: data.GetAllEmployees()[0].Namn})
}

var config Config

func main() {
	readConfig(&config)
	fmt.Println("Database file is: ")
	fmt.Println(config.Database.File)

	data.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	router.GET("/", handleStartPage)
	router.GET("/api/employee", handleGetAllEmployees)
	router.GET("/api/employee/:id", handleGetOneEmployee)

	router.POST("/api/employee", apiEmployeeAdd)
	router.PUT("/api/employee/:id", apiEmployeeUpdateById)

	router.Run(":8080")
}
