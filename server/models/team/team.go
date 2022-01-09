package team

import (
	"apsis/models/employee"
)

type Team struct {
	Id        int
	Employees []employee.Employee
}

func (t *Team) GetTotalScore() (total int) {
	total = 0
	for _, employee := range t.Employees {
		total += employee.Counter
	}
	return
}

func (t *Team) Get(employeeId int) *employee.Employee {
	return &t.Employees[employeeId]
}

func (t *Team) Insert(score int) {
	newEmployee := employee.Employee{
		Id:      t.GetLength(),
		Counter: score,
	}
	t.Employees = append(t.Employees, newEmployee)
}

func (t *Team) Set(employeeId int, newCounter int) {
	t.Employees[employeeId].Counter = newCounter
}

func (t *Team) UpdateTeamScore(newScore int) int {
	diff := t.GetTotalScore() - newScore
	employee := employee.Employee{
		Id:      t.GetLength(),
		Counter: diff,
	}
	t.Employees = append(t.Employees, employee)
	return t.GetTotalScore()
}

func (t *Team) Delete(index int) {
	if len(t.Employees) <= index {
		return
	}

	for i := len(t.Employees) - 1; i >= 0; i-- {
		if index == i {
			t.Employees = append(t.Employees[:i],
				t.Employees[i+1:]...)
		}
	}
	// 0 1 2 3
	for i := index; i <= len(t.Employees)-1; i++ {
		t.Employees[i].Id -= 1
	}
}

func (t *Team) GetCounter(employeeId int) int {
	return t.Employees[employeeId].Counter
}

func (t *Team) SetCounter(employeeId int, newCounter int) {
	t.Employees[employeeId].Counter = newCounter
}

func (t *Team) GetLength() int {
	return len(t.Employees)
}
