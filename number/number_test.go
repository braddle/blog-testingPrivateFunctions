package number_test

import (
	"testing"

	"github.com/braddle/blog-testingPrivateFunctions/number"
)

func TestTwoAddTwoReturnsStringFour(t *testing.T) {
	act := number.Add(2, 2)
	exp := "four"

	assertEquals(t, exp, act)
}

func TestSixAndFiveReturnsNumberNotConvertableString(t *testing.T) {
	act := number.Add(6, 5)
	exp := "Cannot convert 11 to a string"

	assertEquals(t, exp, act)
}

func assertEquals(t *testing.T, exp, act string) {
	if act != exp {
		t.Error("Actual value did not match Expected value")
		t.Logf("Expected: %s", exp)
		t.Logf("Actual:   %s", act)
	}

}