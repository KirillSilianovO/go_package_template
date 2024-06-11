package go_package_template

import "testing"

func TestMultiply(t *testing.T) {
	a := 2
	b := 2

	result := Multiply(a, b)

	if result != 4 {
		t.Errorf("Multiply(%d, %d) = %d; want 4", a, b, result)
	}
}
