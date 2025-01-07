package utils

import "testing"

func TestCleanInput(t *testing.T) {
	input := "  test input  "
	expected := "test input"

	result := CleanInput(input)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"invalid", false},
	}

	for _, test := range tests {
		result := IsIP(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
