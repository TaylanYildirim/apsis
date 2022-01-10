package teams

import (
	"apsis/models/team"
	"reflect"
)

type Teams struct {
	Teams []team.Team
}

func NewTeams() *Teams {
	return &Teams{}
}

func (t Teams) IsEmpty() bool {
	return reflect.DeepEqual(t, Teams{})
}

func (t *Teams) GetTeam(id int) *team.Team {
	if id >= 0 && id < len(t.Teams) {
		return &t.Teams[id]
	}
	return &team.Team{}
}

func (t *Teams) AddNewTeam(newTeam team.Team) {
	t.Teams = append(t.Teams, newTeam)
}

func (t *Teams) GetLength() int {
	return len(t.Teams)
}

func (t *Teams) Delete(i int) {
	if len(t.Teams) == 1 {
		t.Teams = t.Teams[:0]
		return
	}
	t.Teams = append(t.Teams[:i], t.Teams[i+1:]...)
}

func (t *Teams) ListAllTeamScores() *[]map[string]int {
	var res []map[string]int
	for i, team := range t.Teams {
		res = append(res, map[string]int{"TeamID": i, "Score": team.GetTotalScore()})
	}
	return &res
}
