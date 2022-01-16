package employee

import "testing"

func TestEmployee_Get(t *testing.T) {
	tests := []struct {
		expected Employee
		in       int
	}{
		{
			Employee{
				Id:      0,
				Counter: 10,
			},
			10,
		},
		{
			Employee{
				Id:      0,
				Counter: 15,
			},
			15,
		},
	}
	for _, test := range tests {
		if test.expected.Get() != test.in {
			t.Fail()
		}
	}
}

func TestEmployee_Set(t *testing.T) {
	tests := []struct {
		expected Employee
		in       int
	}{
		{
			Employee{
				Id:      0,
				Counter: 10,
			},
			10,
		},
		{
			Employee{
				Id:      0,
				Counter: 15,
			},
			15,
		},
	}
	for _, test := range tests {
		tempEmployee := new(Employee)
		tempEmployee.Set(test.in)
		if tempEmployee.Counter != test.expected.Counter {
			t.Fail()
		}
	}
}