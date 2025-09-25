package employee

import (
	"sync/atomic"
	"time"
)

var lastID atomic.Int64

func init() {
	lastID.Store(100)
}

func GenerateID() int64 {
	return lastID.Add(1)
}

// 1.jadikan public dulu employee
type Employee struct {
	id        int64 //private
	FirstName string
	LastName  string
	HireDate  time.Time
	Salary    float64
}

// 1. constructor return pointer employee
// shareing value
// encapsulation method
func NewEmployee(firstName string, lastName string, hireDate time.Time, salary float64) *Employee {
	return &Employee{
		id:        GenerateID(),
		FirstName: firstName,
		LastName:  lastName,
		HireDate:  hireDate,
		Salary:    salary,
	}
}

// 2. constructor return value employee
func NewEmployeeValue(firstName string, lastName string, hireDate time.Time, salary float64) Employee {
	return Employee{
		id:        GenerateID(),
		FirstName: firstName,
		LastName:  lastName,
		HireDate:  hireDate,
		Salary:    salary,
	}
}

func (e *Employee) SetId(id int64) {
	if e != nil {
		e.id = id
	}
}

func (e *Employee) GetFirstName() string {
	if e != nil {
		return e.FirstName
	}
	return ""
}

func (e *Employee) SetFirstName(firstName string) {
	if e != nil {
		e.FirstName = firstName
	}
}

func (e *Employee) GetLastName() string {
	if e != nil {
		return e.LastName
	}
	return ""
}

func (e *Employee) SetLastName(lastName string) {
	if e != nil {
		e.LastName = lastName
	}
}

// TODO: default value for time.Time is not resolved.
func (e *Employee) GetHireDate() time.Time {
	var ret time.Time
	if e != nil {
		ret = e.HireDate
	}
	return ret
}

func (e *Employee) SetHireDate(hireDate time.Time) {
	if e != nil {
		e.HireDate = hireDate
	}
}

func (e *Employee) GetSalary() float64 {
	if e != nil {
		return e.Salary
	}
	return 0
}

func (e *Employee) SetSalary(salary float64) error {
	if e != nil {
		if salary < minimumWage {
			return ErrInvalidSalaryMin
		} else if salary > maxmumWage {
			return ErrInvalidSalaryMax
		}
		e.Salary = salary
	}
	return nil
}
