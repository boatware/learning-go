package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"Divide 10 by 3", 10.0, 3.0, 3.3333333, false},
	{"Divide 10 by 0", 10.0, 0.0, 0.0, true},
}

func TestDivideTable(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.dividend, tt.divisor)
			if tt.isErr && err == nil {
				t.Error("Expected an error")
			}
			if !tt.isErr && err != nil {
				t.Error(err)
			}
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}
