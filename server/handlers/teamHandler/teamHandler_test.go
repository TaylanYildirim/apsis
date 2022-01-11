package teamHandler

import (
	"apsis/models/employee"
	"apsis/models/team"
	"apsis/models/teams"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewTeam(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/team", nil)
	testTeams := teams.NewTeams()
	w := httptest.NewRecorder()

	CreateNewTeam(testTeams)(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if testTeams.GetLength() == 0 {
		t.Errorf("err: create new team operation failed")
	}
}
func TestDeleteTeam(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/team/%s", "0"), nil)
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

	DeleteTeam(testTeams)(w, req)
	res := w.Result()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if testTeams.GetLength() != 0 {
		t.Errorf("err: team deletion operation failed")
	}
}
func TestGetTeamScore(t *testing.T) {
	type bodyJson struct {
		Data    int    `json:"data"`
		Message string `json:"message"`
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/team/%s", "0"), nil)
	testTeams := teams.NewTeams()
	testTeam := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 10,
		},
			{
				Id:      1,
				Counter: 150,
			}},
	}
	testTeams.AddNewTeam(testTeam)
	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()
	GetTeamScore(testTeams)(w, req)
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(GetTeamScore(testTeams))
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	responseBody := recorder.Body.String()
	var bodyParser bodyJson
	json.Unmarshal([]byte(responseBody), &bodyParser)

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if bodyParser.Data != 160 {
		t.Errorf("get total score test failed")
	}
}
func TestListAllTeamsScores(t *testing.T) {
	type bodyJson struct {
		TeamID int `json:"teamID"`
		Score  int `json:"score"`
	}
	var arr []bodyJson

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/team/%s", "0"), nil)
	testTeams := teams.NewTeams()
	testTeam1 := team.Team{
		Id: 0,
		Employees: []employee.Employee{{
			Id:      0,
			Counter: 10,
		},
			{
				Id:      1,
				Counter: 150,
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
				Counter: 300,
			}},
	}
	testTeams.AddNewTeam(testTeam1)
	testTeams.AddNewTeam(testTeam2)
	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()
	ListAllTeamsScores(testTeams)(w, req)
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(ListAllTeamsScores(testTeams))
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	responseBody := recorder.Body.String()
	json.Unmarshal([]byte(responseBody), &arr)

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	} else if arr[0].Score != 160 {
		t.Errorf("list all team score test failed")
	} else if arr[1].Score != 320 {
		t.Errorf("list all team score test failed")
	}
}