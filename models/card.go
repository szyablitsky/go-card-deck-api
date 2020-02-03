package models

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func NewWithCode(value, suit string) Card {
	code := valueCode(value) + suit[:1]
	return Card{value, suit, code}
}

func valueCode(value string) string {
	switch value {
	case "10":
		return "10"
	default:
		return value[:1]
	}
}
