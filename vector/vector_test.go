package vector

import (
	"testing"
)

func TestIn(t *testing.T) {
	for _, c := range []struct {
		name string
		b    B
		v    V
		want bool
	}{
		{
			name: "Trivial",
			b:    B{V{0, 0}, V{1, 1}},
			v:    V{0, 0},
			want: true,
		},
		{
			name: "Trivial/Outside",
			b:    B{V{0, 0}, V{1, 1}},
			v:    V{100, 100},
			want: false,
		},
		{
			name: "Trivial/Border",
			b:    B{V{0, 0}, V{1, 1}},
			v:    V{1, 1},
			want: false,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := c.b.In(c.v); got != c.want {
				t.Errorf("In() = %v, want = %v", got, c.want)
			}
		})
	}
}
