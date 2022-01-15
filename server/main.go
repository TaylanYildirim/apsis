package main

import (
	"apsis/handlers/employeeHandler"
	"apsis/handlers/teamHandler"
	"apsis/models/teams"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

var router *chi.Mux
var teamsDB = teams.NewTeams()

func init() {
	router = chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func routers() *chi.Mux {
	router.Handle("/*", http.FileServer(http.Dir("./build")))
	router.Post("/team", teamHandler.CreateNewTeam(teamsDB))
	router.Put("/team/{id}", teamHandler.UpdateTeamScore(teamsDB))
	router.Get("/team/{id}", teamHandler.GetTeamScore(teamsDB))
	router.Get("/team/{id}/employees", teamHandler.ListEmployeesScoresInTeam(teamsDB))
	router.Get("/teams", teamHandler.ListAllTeamsScores(teamsDB))
	router.Delete("/team/{id}", teamHandler.DeleteTeam(teamsDB))
	router.Post("/teams/{id}/employee", employeeHandler.CreateNewEmployee(teamsDB))
	router.Delete("/teams/{id}/employees/{eID}", employeeHandler.DeleteEmployee(teamsDB))
	router.Get("/employees", employeeHandler.ListAllEmployees(teamsDB))
	router.Put("/employees/{eID}", employeeHandler.EditEmployee(teamsDB))
	return router

}
func main() {
	routers()
	port := os.Getenv("PORT")
	defaultPort := "10000"

	if !(port == "") {
		log.Fatal(http.ListenAndServe(":"+port, router))
	} else {
		log.Printf("Starting up on http://localhost:%s", defaultPort)
		log.Fatal(http.ListenAndServe(":"+defaultPort, router))
	}
}
