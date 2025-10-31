package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Members []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

var (
	ErrUnknownTeam = errors.New("unknown team")
	ErrDraw        = errors.New("draw not counted")
)

func (l *League) MatchResult(team1 string, scoreT1 int, team2 string, scoreT2 int) error {
	// Guardchecks, Sanitycheck if Team is in the map already
	if _, ok := l.Teams[team1]; !ok {
		return fmt.Errorf("%w: %s", ErrUnknownTeam, team1)
	}
	if _, ok := l.Teams[team2]; !ok {
		return fmt.Errorf("%w: %s", ErrUnknownTeam, team2)
	}
	if scoreT1 == scoreT2 {
		return ErrDraw
	}
	if scoreT1 > scoreT2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
	return nil
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Members: []string{"Member1", "Member2", "Member3", "Member4", "Member5"},
			},
			"France": {
				Name:    "France",
				Members: []string{"Member1", "Member2", "Member3", "Member4", "Member5"},
			},
			"India": {
				Name:    "India",
				Members: []string{"Member1", "Member2", "Member3", "Member4", "Member5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Members: []string{"Member1", "Member2", "Member3", "Member4", "Member5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	RankPrinter(l, os.Stdout)
}
