package math

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{5, 6, 11},
	}
	for _, c := range cases {
		got := Sum(c.a, c.b)
		if got != c.want {
			t.Errorf("Sum(%d,%d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
