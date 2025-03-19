package main

import "testing"

func TestCamelCaseToTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"luxuryApartments", "Luxury Apartments"},
		{"rentOffice", "Rent Office"},
		{"lowCostHousing", "Low Cost Housing"},
		{"pawnShop", "Pawn Shop"},
		{"zMart", "Z Mart"},
		{"monolithBurgers", "Monolith Burgers"},
		{"qtClothing", "Qt Clothing"},
		{"socketCity", "Socket City"},
		{"hiTechU", "Hi Tech U"},
		{"employmentOffice", "Employment Office"},
		{"factory", "Factory"},
		{"bank", "Bank"},
		{"blacksMarket", "Blacks Market"},
	}

	for _, test := range tests {
		result := camelCaseToTitle(test.input)
		if result != test.expected {
			t.Errorf("camelCaseToTitle(%q) WANT: %q, GOT: %q;", test.input, test.expected, result)
		}
	}
}
