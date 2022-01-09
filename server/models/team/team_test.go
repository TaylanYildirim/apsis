package team

import (
	"apsis/models/employee"
	"testing"
)

func TestTeam_Insert(t *testing.T) {
	employee1, employee2, employee3 := new(employee.Employee), new(employee.Employee), new(employee.Employee)
	employee1.Counter = 10
	employee2.Counter = 40
	employee3.Counter = 100

	testEmployee := make([]employee.Employee, 5)
	for i := range testEmployee {
		testEmployee[i] = employee.Employee{}
	}

	testEmployee2 := make([]employee.Employee, 10)
	for i := range testEmployee {
		testEmployee[i] = employee.Employee{}
	}

	tests := []struct {
		in       Team
		expected int
	}{
		{
			*new(Team),
			1,
		},
		{
			Team{
				Employees: testEmployee,
			},
			6,
		},
		{
			Team{
				Employees: testEmployee2,
			},
			11,
		},
	}
	for _, test := range tests {
		test.in.Insert(10)
		if len(test.in.Employees) != test.expected {
			t.Fail()
		}
	}
}
func TestTeam_Delete(t *testing.T) {
	teamsArr1 := make([]employee.Employee, 3)
	for i := range teamsArr1 {
		teamsArr1[i] = employee.Employee{}
	}
	teamsArr2 := make([]employee.Employee, 10)
	for i := range teamsArr1 {
		teamsArr1[i] = employee.Employee{}
	}
	tests := []struct {
		in       int
		expected int
	}{
		{
			3,
			2,
		},
		{
			10,
			9,
		},
	}
	teams := []Team{
		{
			Employees: teamsArr1,
		},
		{
			Employees: teamsArr2,
		},
	}

	for i, test := range tests {
		teams[i].Delete(test.in - 1)
		if teams[i].GetLength() != test.expected {
			t.Fail()
		}
	}
}

func TestTeam_GetCounter(t *testing.T) {
	tests := []struct {
		in       Team
		expected int
	}{
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 10}},
			},
			10,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 100}},
			},
			100,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: -10}},
			},
			-10,
		},
	}
	for _, test := range tests {
		if test.in.GetCounter(test.in.Id) != test.expected {
			t.Fail()
		}
	}
}

func TestTeam_GetLength(t *testing.T) {
	tests := []struct {
		in       Team
		expected int
	}{
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 10}},
			},
			1,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 1}, {Counter: 2}, {Counter: 5}},
			},
			3,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{},
			},
			0,
		},
	}
	for _, test := range tests {
		if test.in.GetLength() != test.expected {
			t.Fail()
		}
	}
}

func TestTeam_Get(t *testing.T) {
	tests := []struct {
		in       Team
		expected []int
	}{
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Id: 0, Counter: 10}},
			},
			[]int{0},
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Id: 0, Counter: 1}, {Id: 1, Counter: 2}, {Id: 2, Counter: 5}},
			},
			[]int{0, 1, 2},
		},
	}
	for i, test := range tests {
		if test.in.Get(i).Id != test.expected[i] {
			t.Fail()
		}
	}
}

func TestTeam_GetTotalScore(t *testing.T) {
	tests := []struct {
		in       Team
		expected int
	}{
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 10}},
			},
			10,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: 1}, {Counter: 2}, {Counter: 5}},
			},
			8,
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Counter: -1}, {Counter: -2}, {Counter: -5}},
			},
			-8,
		},
	}
	for _, test := range tests {
		if test.in.GetTotalScore() != test.expected {
			t.Fail()
		}
	}
}

func TestTeam_Set(t *testing.T) {
	tests := []struct {
		in       Team
		expected []int
	}{
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Id: 0, Counter: 10}},
			},
			[]int{0},
		},
		{
			Team{
				Id:        0,
				Employees: []employee.Employee{{Id: 0, Counter: 1}, {Id: 1, Counter: 2}, {Id: 2, Counter: 5}},
			},
			[]int{0, 1, 2},
		},
	}
	for i, test := range tests {
		test.in.Set(i, i)
		if test.in.Employees[i].Counter != test.expected[i] {
			t.Fail()
		}
	}
}

func TestTeam_SetCounter(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{
			10,
			10,
		},
		{
			8,
			8,
		},
	}
	teamsArr1 := make([]employee.Employee, 3)
	for i := range teamsArr1 {
		teamsArr1[i] = employee.Employee{}
	}
	teamsArr2 := make([]employee.Employee, 10)
	for i := range teamsArr1 {
		teamsArr1[i] = employee.Employee{}
	}

	teams := []Team{
		{
			Employees: teamsArr1,
		},
		{
			Employees: teamsArr2,
		},
	}

	for i, test := range tests {
		teams[0].SetCounter(i, i)
		if teams[0].Employees[0].Counter == test.expected {
			t.Fail()
		}
	}
}
