package requests

import "testing"

func TestGetLocales(t *testing.T) {
	locales, err := GetLocales()

	if !(locales.Status.IsValid() && locales.Status.IsComplete()) {
		t.Fatalf("Status is not valid: %v", locales.Status)
	}

	if len(locales.Locales) == 0 {
		t.Fatalf("No locales returned")
	}

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
