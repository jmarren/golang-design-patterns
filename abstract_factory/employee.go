package abstractfactory

type IEmployee interface {
	getLocation() string
	getSalary() int
	setLocation(string)
	setSalary(int)
}

type Employee struct {
	location string
	salary   int
}

func (e *Employee) setLocation(location string) {
	e.location = location
}

func (e *Employee) setSalary(salary int) {
	e.salary = salary
}

func (e *Employee) getSalary() int {
	return e.salary
}

func (e *Employee) getLocation() string {
	return e.location
}
