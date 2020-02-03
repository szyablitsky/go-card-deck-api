package models

import "testing"

func TestNewWithCode(t *testing.T) {
	value := "VALUE"
	suit := "SUIT"
	wantCode := "VS"

	got := NewWithCode(value, suit)

	if got.Value != value {
		t.Errorf("NewWithCode(%q, %q).Value == %q, got %q", value, suit, value, got.Value)
	}
	if got.Suit != suit {
		t.Errorf("NewWithCode(%q, %q).Suit == %q, got %q", value, suit, suit, got.Suit)
	}
	if got.Code != wantCode {
		t.Errorf("NewWithCode(%q, %q).Code == %q, got %q", value, suit, wantCode, got.Code)
	}
}

func TestValueCode(t *testing.T) {
	want := "10"
	if got := valueCode("10"); got != want {
		t.Errorf("valueCode(%q) should return %q unmodified", want, want)
	}

	input := "VALUE"
	want = "V"
	if got := valueCode(input); got != want {
		t.Errorf("valueCode(%q) == %q, want %q", input, got, want)
	}
}
