// math/math_test.go:
package math_test

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
	if math.NthPrime(450001) != 6581969 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(24367) != 279119 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(24370) != 279131 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(6950000) != 122017249 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(8500000) != 151048969 {
		t.Fatal("unexpected prime value")
	}
	if math.NthPrime(3) != 5 {
		t.Fatal("unexpected prime value")
	}
	//if math.NthPrime(142) != 821 {
	//	t.Fatal("unexpected prime value")
}
