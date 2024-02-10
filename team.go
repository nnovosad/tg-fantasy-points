package main

import (
	"fmt"
	"sort"
)

type PlayersSlice []PlayerInfo

func (p PlayersSlice) Less(i, j int) bool {
	if p[i].IsStarting && !p[j].IsStarting {
		return true
	} else if !p[i].IsStarting && p[j].IsStarting {
		return false
	}

	roleOrder := map[string]int{
		"GOALKEEPER": 1,
		"DEFENDER":   2,
		"MIDFIELDER": 3,
		"FORWARD":    4,
	}

	return roleOrder[p[i].SeasonPlayer.Role] < roleOrder[p[j].SeasonPlayer.Role]
}

func (p PlayersSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PlayersSlice) Len() int      { return len(p) }

func displayTeam(response Response) {
	squad := response.Data["id_105467854"].Squads[0]
	players := squad.CurrentTourInfo.Players

	sort.Sort(PlayersSlice(players))

	for _, player := range squad.CurrentTourInfo.Players {
		fmt.Printf(
			"%v (%v). %v. Price: %v. Is starting: %v. Is Captain: %v. Is Vice Captain: %v. \n",
			player.SeasonPlayer.Name,
			player.SeasonPlayer.Team.Name,
			player.SeasonPlayer.Role,
			player.SeasonPlayer.Price,
			player.IsStarting,
			player.IsCaptain,
			player.IsViceCaptain,
		)
	}
}
