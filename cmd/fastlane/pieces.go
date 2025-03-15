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

type location struct {
	img              string
	name             string
	relativeDistance int
	pos              int
}

type Accommodator interface {
	HowMuch() int
	PayRent(payment int) int
	Rest(time int) int
	Talk(event string) string
	Evict(p player) int
}

type Employer interface {
	Apply(edu, exp int) bool
	Work(time int) int
	Promote(p player) int
	Quit()
	Talk(event string) string
}

type Seller[T any] interface {
	Buy(money int) T
	Talk(event string) string
}

type Buyer[T any] interface {
	Sell(item T) int
	Talk(event string) string
}

type gameMap struct {
	luxuryApartments location
	rentOffice       location
	lowCostHousing   location
	pawnShop         location
	zMart            location
	monolith         location
	qtClothing       location
	socketCity       location
	hiTechU          location
	employmentOffice location
	factory          location
	bank             location
	blacksMarket     location
}

const (
	george name = "George"
	john   name = "John"
	paul   name = "Paul"
	ringo  name = "Ringo"
)

type player int

const (
	playerCountKey   = "playerCount"
	educationGoalKey = "educationGoalKey"
	moneyGoalKey     = "moneyGoalKey"
	happinessGoalKey = "happinessGoalKey"
	careerGoalKey    = "careerGoalKey"
)
