package employeeHandler

import (
	"apsis/handlers"
	"apsis/models/employee"
	"apsis/models/team"
	"apsis/models/teams"
	"apsis/util/stringUtil"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func CreateNewEmployee(teamsDB *teams.Teams) http.HandlerFunc {
	type body struct {
		Score int `json:"score"`
	}
	fn := func(w http.ResponseWriter, r *http.Request) {
		var body body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			handlers.SetHTTPStatus(w, http.StatusNotFound, "StatusNotFound", 0)
			return
		}
		teamId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		if teamsDB.GetTeam(teamId).GetLength() == 0 {
			teamsDB.AddNewTeam(team.Team{
				Id: teamId,
				Employees: []employee.Employee{{
					Id:      0,
					Counter: body.Score,
				}},
			})
		} else {
			teamsDB.GetTeam(teamId).Insert(body.Score)
		}
		handlers.SetHTTPStatus(w, http.StatusOK, "created", 200)
		return
	}
	return fn
}

func DeleteEmployee(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		teamId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		employeeId := stringUtil.ParseUint(chi.URLParam(r, "eID"))

		teamsDB.GetTeam(teamId).Delete(employeeId)
		for _, team := range teamsDB.Teams {
			if teamsDB.GetTeam(team.Id).GetLength() == 0 {
				teamsDB.Delete(team.Id)
			}
		}
		handlers.SetHTTPStatus(w, http.StatusOK, fmt.Sprintf("Employee id: %d deleted", employeeId), 200)
		return
	}
	return fn
}

func ListAllEmployees(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var allEmployees []map[string]int
		for _, team := range teamsDB.Teams {
			teamScore := team.GetTotalScore()
			for _, employee := range team.Employees {
				allEmployees = append(allEmployees, map[string]int{
					"TeamID":     team.Id,
					"EmployeeID": employee.Id,
					"Score":      employee.Counter,
					"TeamScore":  teamScore,
				})
			}
		}
		handlers.SetHTTPResponse(w, http.StatusOK, &allEmployees)
		return
	}
	return fn
}

func EditEmployee(teamsDB *teams.Teams) http.HandlerFunc {
	type body struct {
		Score     int `json:"score"`
		TeamScore int `json:"teamScore"`
		TeamId    int `json:"teamId"`
	}
	fn := func(w http.ResponseWriter, r *http.Request) {
		var body body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			handlers.SetHTTPStatus(w, http.StatusNotFound, "StatusNotFound", 0)
			return
		}
		employeeId := stringUtil.ParseUint(chi.URLParam(r, "eID"))
		get := teamsDB.GetTeam(body.TeamId).Get(employeeId)
		get.Set(body.Score)
		handlers.SetHTTPStatus(w, http.StatusOK, "OKv", 200)
		return
	}
	return fn
}
