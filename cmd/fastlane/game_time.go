package main

import "math"

// Get the travel time from one location
// to the next.
func (gm GameModel) CalcTravelTime(from, to *location) float64 {
	return math.Abs(float64(from.pos-to.pos)) * travelTimeCost
}

// The hours that the player has remaining in a
// turn.
func (gm *GameModel) Hour() float64 {
	return gm.Hours
}

// Used to decrement the player's hours, based
// on either an event, or their own actions.
func (gm *GameModel) DecrementHours(t float64) {
	gm.Hours -= t
}
