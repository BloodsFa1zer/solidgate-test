package model

type Card struct {
	Number         string         `json:"card_number"`
	ExpirationDate ExpirationDate `json:"expiration_date"`
}

type ExpirationDate struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}
