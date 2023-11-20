package tests

import (
	"testing"

	"github.com/nabinbhusal80/goBoilerplate/internal/utilities"
)

func TestCalculations(t *testing.T) {
	result := utilities.CalcAdd(1, 2)
	if result != 3 {
		t.Error("Expected 3")
	}
	t.Log("Result passed, result is: ", result)
}
