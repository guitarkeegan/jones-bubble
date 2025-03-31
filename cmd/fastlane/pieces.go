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
	hours            float64
	week             int
	month            int
}

const (
	// characters
	george name = "George"
	john   name = "John"
	paul   name = "Paul"
	ringo  name = "Ringo"

	// game stats
	relaxationDefault int = 10
	increaseRest      int = 3

	// time trackers
	hoursDefault float64 = 60
	weekDefault  int     = 1
	monthDefault int     = 1

	// hour costs
	starvationTimeCost          = 20.0
	doctorVisitTimeCost         = 10.0
	workingTimeCost             = 6.0
	relaxingTimeCost            = 6.0
	studyingTimeCost            = 6.0
	applyForLoanTimeCost        = 4.0
	applyingForJobTimeCost      = 4.0
	enterTimeCost               = 2.0
	visitBrokerTimeCost         = 2.0
	purchasingNewspaperTimeCost = 1.0
	travelTimeCost              = 0.625
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
