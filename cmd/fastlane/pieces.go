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
	relaxation       int
	hours            int
	week             int
	month            int
}

const (
	george name = "George"
	john   name = "John"
	paul   name = "Paul"
	ringo  name = "Ringo"

	relaxationDefault int = 10
	increaseRest      int = 3

	hoursDefault int = 60
	weekDefault  int = 1
	monthDefault int = 1
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
	lowCostHousingKey   = "lowCostHousingKey"
	rentOfficeKey       = "rentOfficeKey"
)
