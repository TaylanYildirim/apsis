package teams

import (
	"apsis/models/employee"
	"apsis/models/team"
	"reflect"
	"testing"
)

func TestNewTeams(t *testing.T) {
	if !NewTeams().IsEmpty() {
		t.Fail()
	}
}

func TestTeams_GetLength(t *testing.T) {
	testTeams := NewTeams()
	tests := []struct {
		in       int
		expected int
	}{
		{
			3,
			3,
		},
		{
			-1,
			0,
		},
		{
			10,
			10,
		},
	}
	for _, test := range tests {
		for i := 0; i < test.in; i++ {
			testTeams.AddNewTeam(team.Team{})
		}
		if testTeams.GetLength() != len(testTeams.Teams) {
			t.Fail()
		}
	}
}

func TestTeams_Delete(t *testing.T) {
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
	for _, test := range tests {
		testTeams := NewTeams()
		for i := 0; i < test.in; i++ {
			testTeams.AddNewTeam(team.Team{})
		}
		testTeams.Delete(0)
		if testTeams.GetLength() != test.expected {
			t.Fail()
		}
	}
}
func TestTeams_GetTeam(t *testing.T) {
	//	s0, s1, s2 := new(Foo), new(Foo), new(Foo)
	var employeeArr []employee.Employee
	var teamArr []team.Team
	testTeams := NewTeams()
	testEmployees := make([]*employee.Employee, 5)
	testTeam := make([]*team.Team, 5)
	for i := range testTeam {
		testEmployees[i] = &employee.Employee{}
		employeeArr = append(employeeArr, *testEmployees[i])
		testTeam[i] = &team.Team{
			Id:        i,
			Employees: employeeArr,
		}
		teamArr = append(teamArr, *testTeam[i])
		testTeams.AddNewTeam(*testTeam[i])
	}

	tests := []struct {
		in       int
		expected int
	}{
		{
			3,
			3,
		},
	}
	for _, test := range tests {
		for i := 0; i < test.in; i++ {
			if reflect.DeepEqual(teamArr[i], testTeams.GetTeam(i)) {
				t.Fail()
			}
		}
	}
}

func TestTeams_IsEmpty(t *testing.T) {
	team1, team2, team3 := new(team.Team), new(team.Team), new(team.Team)
	testTeams := NewTeams()
	testTeams.AddNewTeam(*team1)
	testTeams.AddNewTeam(*team2)
	testTeams.AddNewTeam(*team3)
	tests := []struct {
		in       int
		expected bool
	}{
		{
			3,
			false,
		},
		{
			6,
			false,
		},
		{
			0,
			true,
		},
	}
	for _, test := range tests {
		testTeams = &Teams{}
		for i := 0; i < test.in; i++ {
			testTeams.AddNewTeam(*team1)
		}
		if testTeams.IsEmpty() != test.expected {
			t.Fail()
		}
	}
}

func TestTeams_ListAllTeamScores(t *testing.T) {
	testTeam := make([]team.Team, 5)
	for i := range testTeam {
		testTeam[i] = team.Team{}
	}

	tests := []struct {
		in       Teams
		expected int
	}{
		{
			*new(Teams),
			0,
		},
		{
			Teams{
				Teams: testTeam,
			},
			5,
		},
	}
	for _, test := range tests {
		if len(*test.in.ListAllTeamScores()) != test.expected {
			t.Fail()
		}
	}
}

func TestTeams_AddNewTeam(t *testing.T) {
	testTeam := make([]team.Team, 5)
	for i := range testTeam {
		testTeam[i] = team.Team{}
	}

	tests := []struct {
		in       Teams
		expected int
	}{
		{
			*new(Teams),
			1,
		},
		{
			Teams{
				Teams: testTeam,
			},
			6,
		},
	}
	for _, test := range tests {
		test.in.AddNewTeam(team.Team{})
		if len(test.in.Teams) != test.expected {
			t.Fail()
		}
	}
}
