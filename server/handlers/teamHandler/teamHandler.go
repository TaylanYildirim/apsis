package teamHandler

import (
	"apsis/handlers"
	"apsis/models/employee"
	"apsis/models/team"
	"apsis/models/teams"
	"apsis/util/stringUtil"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func CreateNewTeam(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		newTeam := team.Team{
			Id:        teamsDB.GetLength(),
			Employees: []employee.Employee{},
		}
		teamsDB.AddNewTeam(newTeam)
	}
	return fn
}

func GetTeamScore(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		teamScore := teamsDB.GetTeam(stringUtil.ParseUint(id)).GetTotalScore()
		handlers.SetHTTPStatus(w, http.StatusOK, "OK", teamScore)
		return
	}
	return fn
}

func ListEmployeesScoresInTeam(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		teamId := chi.URLParam(r, "id")
		employeesScore := teamsDB.GetTeam(stringUtil.ParseUint(teamId)).ListEmployeesScoreInTeam()
		handlers.SetHTTPResponse(w, http.StatusOK, employeesScore)
		return
	}
	return fn
}

func UpdateTeamScore(teamsDB *teams.Teams) http.HandlerFunc {
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
		id := chi.URLParam(r, "id")
		teamScore := teamsDB.GetTeam(stringUtil.ParseUint(id)).UpdateTeamScore(body.Score)
		handlers.SetHTTPStatus(w, http.StatusOK, "OK", teamScore)
		return
	}
	return fn
}

func ListAllTeamsScores(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		teamScores := teamsDB.ListAllTeamScores()
		handlers.SetHTTPResponse(w, http.StatusOK, teamScores)
		return
	}
	return fn
}

func DeleteTeam(teamsDB *teams.Teams) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		teamsDB.Delete(stringUtil.ParseUint(id))
		handlers.SetHTTPStatus(w, http.StatusOK, "Ok", stringUtil.ParseUint(id))
		return
	}
	return fn
}
