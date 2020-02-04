package repository

import "testing"

func TestValidateCards(t *testing.T) {
	input := []string{"AC", "TS"}
	got := ValidateCards(input)
	if got == false {
		t.Errorf("ValidateCards(%v) == true, got %v", input, got)
	}

	input = []string{"WRONG"}
	got = ValidateCards(input)
	if got == true {
		t.Errorf("ValidateCards(%v) == false, got %v", input, got)
	}
}
