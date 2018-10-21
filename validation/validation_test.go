package validation

import (
	"testing"
)

func TestRequired(t *testing.T) {
	validation := New()
	required := validation.Validate("Fooo", "required")

	if required != nil {
		t.Fatal(required)
	}
}

func TestInteger(t *testing.T) {
	validation := New()
	integer := validation.Validate("50", "integer")

	if integer != nil {
		t.Fatal(integer)
	}
}

func TestNumeric(t *testing.T) {
	validation := New()
	numeric := validation.Validate(50.5, "numeric")

	if numeric != nil {
		t.Fatal(numeric)
	}
}

func TestMax(t *testing.T) {
	validation := New()
	max := validation.Validate("Rachid", "max:2")

	if max != nil {
		t.Fatal(max)
	}
}

func TestMin(t *testing.T) {
	validation := New()
	min := validation.Validate(51, "min:50")

	if min != nil {
		t.Fatal(min)
	}
}

func TestBetween(t *testing.T) {
	validation := New()
	between := validation.Validate(0, "between:50.1,60|required")

	if between != nil {
		t.Fatal(between)
	}
}

func TestEmail(t *testing.T) {
	validation := New()
	email := validation.Validate("lafriakh.rachid@gmail.com", "email")

	if email != nil {
		t.Fatal(email)
	}
}
