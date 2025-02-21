package data

type Employee struct {
	Id     int
	Age    int
	City   string
	Namn   string
	Age2   int
	Active bool
}

func (emp Employee) CalculateSalary() int {
	if emp.Namn == "Stefan" {
		return 1000
	}
	return 10
}

func CalculateSalary(emp Employee) int {
	if emp.Namn == "Stefan" {
		return 1000
	}
	return 10
}
