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
}
