package mul

import "testing"

var mul_tests = []struct {
	a, b      int
	expected  int
}{
	{1, 1, 1},
	{1, 2, 2},
}

func TestMul(t *testing.T) {
	for _, mt := range mul_tests {
		if v := Mul(mt.a, mt.b); v != mt.expected {
			t.Errorf("Mul(%d, %d) returned %d, expected %d", mt.a, mt.b, v, mt.expected)
		}
	}
}
