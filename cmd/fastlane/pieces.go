package main

type name = string

type character struct {
	name             name
	currentHappiness int
	happinessGoal    int
	currentMoney     int
	moneyGoal        int
	currentEducation int
	educationGoal    int
	currentCareer    int
	careerGoal       int
}

const (
	george name = "George"
	john   name = "John"
	paul   name = "Paul"
	ringo  name = "Ringo"
)

type player int

const (
	playerCountKey      = "playerCount"
	educationGoalKey    = "educationGoalKey"
	moneyGoalKey        = "moneyGoalKey"
	happinessGoalKey    = "happinessGoalKey"
	careerGoalKey       = "careerGoalKey"
	destinationKey      = "destinationKey"
	luxuryApartmentsKey = "luxuryApartmentsKey"
)
