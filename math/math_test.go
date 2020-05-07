// math/math_test.go:
package main_test

import (
	math "hsecode.com/stdlib/math"
	"testing"
)

func TestCompare(t *testing.T) {
	if math.NthPrime(3) != 5 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(141) != 811 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(3) != 5 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(142) != 821 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(3) != 5 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(24368) != 279121 {
		t.Fatal("unexpected prime value")
	}
}
