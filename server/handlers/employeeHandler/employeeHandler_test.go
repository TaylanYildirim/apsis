package employeeHandler

import (
	"apsis/models/employee"
	"apsis/models/team"
	"apsis/models/teams"
	"apsis/util/errorUtil"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewEmployee(t *testing.T) {
	reqBody := map[string]int{
		"score": 5,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/teams/%s/employee", "0"), bytes.NewReader(body))
	var m map[string]int
	err := json.NewDecoder(req.Body).Decode(&m)
	errorUtil.Check(err)
	req.Body.Close()
	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}
}

func TestDeleteEmployee(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/teams/%s/employees/%s", "0", "0"), nil)
	testTeams := teams.NewTeams()
	testTeam := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 10,
		},
			{
				Id:      1,
				Counter: 15,
			}},
	}
	testTeams.AddNewTeam(testTeam)

	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()

	DeleteEmployee(testTeams)(w, req)

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if teamsDBLength := testTeams.GetLength(); teamsDBLength != 1 {
		t.Errorf("err: delete operation failed")
	}
}
func TestEditEmployee(t *testing.T) {
	reqBody := map[string]int{
		"score":     5,
		"teamScore": 5,
		"teamId":    0,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/employees/%s", "0"), bytes.NewReader(body))
	testTeams := teams.NewTeams()
	testTeam := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 10,
		},
			{
				Id:      1,
				Counter: 15,
			}},
	}
	testTeams.AddNewTeam(testTeam)

	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()

	EditEmployee(testTeams)(w, req)

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if testTeams.GetTeam(0).GetCounter(0) != 5 {
		t.Errorf("err: edit operation failed")
	}
}

func TestListAllEmployees(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/employees", nil)

	testTeams := teams.NewTeams()
	testTeam1 := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 10,
		},
			{
				Id:      1,
				Counter: 15,
			}},
	}
	testTeam2 := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 20,
		},
			{
				Id:      1,
				Counter: 25,
			}},
	}
	testTeams.AddNewTeam(testTeam1)
	testTeams.AddNewTeam(testTeam2)
	w := httptest.NewRecorder()

	ListAllEmployees(testTeams)(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}
}
