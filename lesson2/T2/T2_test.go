// T2_test.go
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	Add := Fac("+")
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	Mu := Fac("*")
	result := Mu(4, 5)
	expected := 20

	if result != expected {
		t.Errorf("Mu(4, 5) = %d; expected %d", result, expected)
	}
}
