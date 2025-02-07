package flaghelpers

import (
	"github.com/concourse/concourse/go-concourse/concourse"
	"strings"

	"github.com/concourse/concourse/fly/rc"
	"github.com/jessevdk/go-flags"
)

type TeamFlag string

func (flag *TeamFlag) Complete(match string) []flags.Completion {
	fly := parseFlags()

	target, err := rc.LoadTarget(fly.Target, false)
	if err != nil {
		return []flags.Completion{}
	}

	teams, err := target.Client().ListTeams()
	if err != nil {
		return []flags.Completion{}
	}

	comps := []flags.Completion{}
	for _, team := range teams {
		if strings.HasPrefix(team.Name, match) {
			comps = append(comps, flags.Completion{Item: team.Name})
		}
	}

	return comps
}

func (flag TeamFlag) Name() string {
	return string(flag)
}

func (flag TeamFlag) LoadTeam(target rc.Target) (concourse.Team, error) {
	team := target.Team()
	var err error
	if flag.Name() != "" {
		team, err = target.FindTeam(flag.Name())
		if err != nil {
			return nil, err
		}
	}
	return team, nil
}
