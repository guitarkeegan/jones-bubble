package main

import (
	"testing"

	"github.com/charmbracelet/huh"
)

func TestCalcTravelTime(t *testing.T) {

	gm := GameModel{
		Board:          map[string]*location{},
		MainMap:        "",
		GameState:      "",
		ActionsMenu:    &huh.Form{},
		CurrentLoc:     &location{},
		GameMsg:        "",
		Hours:          0,
		GameMsgCounter: 0,
	}

	locs := []location{
		{pos: 0},
		{pos: 3},
		{pos: 3},
		{pos: 3},
		{pos: 5},
		{pos: 3},
		{pos: 3},
		{pos: 5},
	}

	tests := []struct {
		title  string
		start  *location
		end    *location
		expect float64
	}{
		{
			title:  "pos 0 - 3",
			start:  &locs[0],
			end:    &locs[1],
			expect: 1.875,
		},
		{
			title:  "pos 3 - 3 (same pos)",
			start:  &locs[1],
			end:    &locs[2],
			expect: 0,
		},
		{
			title:  "pos 5 - 3",
			start:  &locs[3],
			end:    &locs[4],
			expect: 1.25,
		},
		{
			title:  "pos 3 - 5 (backward)",
			start:  &locs[4],
			end:    &locs[5],
			expect: 1.25,
		},
	}

	for _, test := range tests[1:] {
		res := gm.CalcTravelTime(test.start, test.end)
		if res != test.expect {
			t.Errorf("%s, Got: %f, Want: %f", test.title, res, test.expect)
		}
	}
	// 0.625

}
