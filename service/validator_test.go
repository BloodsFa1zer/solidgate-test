package service

import (
	"errors"
	"solidgate-test/model"
	"solidgate-test/util"
	"testing"
	"time"
)

func TestIsCardLuhnAlgorithmValid(t *testing.T) {
	cv := &CardValidator{}
	tests := []struct {
		name       string
		cardNumber string
		expected   bool
	}{
		// reference for data taken https://www.getcreditcardnumbers.com/
		{"Valid Visa 13", "4929790773754", true},
		{"Invalid Visa 13", "4111111111112", false},
		{"Valid Visa 16", "4532015112830366", true},
		{"Invalid Visa 16", "4532015112830367", false},
		{"Valid Mastercard", "5555555555554444", true},
		{"Invalid Mastercard", "5555555555554445", false},
		{"Valid Discover", "6011111111111117", true},
		{"Invalid Discover", "6011111111111118", false},
		{"Valid American Express", "378282246310005", true},
		{"Invalid American Express", "378282246310006", false},
		{"Valid Diner's Club", "30569309025904", true},
		{"Invalid Diner's Club", "30569309025905", false},
		{"Valid enRoute", "201424691322712", true},
		{"Invalid enRoute", "201424691322710", false},
		{"Invalid Maestro", "500000000000", false},
		{"Valid JCB 15", "180000000000002", true},
		{"Valid JCB 16", "3530111333300000", true},
		{"Invalid JCB", "3530111333300001", false},
		{"Valid China UnionPay", "6200000000000005", true},
		{"Invalid China UnionPay", "6200000000000006", false},
		{"Dummy data", "dfkopdokvdpo", false},
		{"Less numbers than needed", "123467509", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cv.isCardLuhnAlgorithmValid(tt.cardNumber); got != tt.expected {
				t.Errorf("isCardLuhnAlgorithmValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsCardDateValid(t *testing.T) {
	cv := &CardValidator{}
	currentYear, currentMonth, _ := time.Now().Date()
	tests := []struct {
		name  string
		year  int
		month int
		valid bool
		error *util.Error
	}{
		{"Valid future date", currentYear + 1, int(currentMonth), true, nil},
		{"Valid current month", currentYear, int(currentMonth), true, nil},
		{"Invalid past year", currentYear - 1, int(currentMonth), false, util.ExpiredCardError()},
		{"Invalid past month", currentYear, int(currentMonth) - 1, false, util.ExpiredCardError()},
		{"Invalid month 13 of 12", currentYear, 13, false, util.InvalidCardNumberError()},
		{"Invalid month -1 of 12", currentYear, 13, false, util.InvalidCardNumberError()},
		{"Valid last month of year", currentYear, 12, true, nil},
		{"Valid first month of next year", currentYear + 1, 1, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := cv.isCardDateValid(tt.year, tt.month)
			if valid != tt.valid {
				t.Errorf("Expected validity %v, got %v", tt.valid, valid)
			}
			if tt.error != nil && err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err != nil && tt.error == nil {
				t.Errorf("Expected error %v, got no error", err)
			}
		})
	}
}

func TestIsValidCard(t *testing.T) {
	cv := &CardValidator{}
	currentYear, currentMonth, _ := time.Now().Date()

	validCard := model.Card{
		Number: "4532015112830366",
		ExpirationDate: struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		}{
			Month: int(currentMonth) + 1,
			Year:  currentYear,
		},
	}

	tests := []struct {
		name     string
		card     model.Card
		expected bool
		error    *util.Error
	}{
		{
			name:     "Valid card",
			card:     validCard,
			expected: true,
			error:    nil,
		},
		{
			name: "Invalid card number",
			card: model.Card{
				Number:         "1234567890123456", // Invalid Luhn
				ExpirationDate: validCard.ExpirationDate,
			},
			expected: false,
			error:    util.InvalidCardNumberError(),
		},
		{
			name: "Expired card",
			card: model.Card{
				Number: validCard.Number,
				ExpirationDate: struct {
					Month int `json:"month"`
					Year  int `json:"year"`
				}{
					Month: int(currentMonth) - 1,
					Year:  currentYear,
				},
			},
			expected: false,
			error:    util.ExpiredCardError(),
		},
		{
			name: "Invalid expiration date",
			card: model.Card{
				Number: validCard.Number,
				ExpirationDate: struct {
					Month int `json:"month"`
					Year  int `json:"year"`
				}{
					Month: 13,
					Year:  currentYear,
				},
			},
			expected: false,
			error:    util.InvalidExpirationDateError(),
		},
		{
			name: "Invalid card number and expiration date",
			card: model.Card{
				Number: "1234567890123456",
				ExpirationDate: struct {
					Month int `json:"month"`
					Year  int `json:"year"`
				}{
					Month: 13,
					Year:  currentYear,
				},
			},
			expected: false,
			error:    util.InvalidCardNumberAndExpirationDateError(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cv.IsValidCard(tt.card)
			//if valid != tt.expected {
			//	t.Errorf("Expected validity %v, got %v", tt.expected, valid)
			//}
			if (err == nil && tt.error != nil) || (err != nil && tt.error == nil) || (err != nil && tt.error != nil && errors.Is(err, tt.error)) {
				t.Errorf("Expected error %v, got %v", tt.error, err)
			}
		})
	}
}
