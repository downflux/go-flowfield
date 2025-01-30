package vector

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v := V{1, 1}
	for _, c := range []struct {
		name string
		v    V
		u    V
		buf  V
		want V
	}{
		{
			name: "Trivial",
			v:    V{1, 1},
			u:    V{10, 11},
			buf:  V{},
			want: V{11, 12},
		},
		{
			name: "OverwriteBuffer",
			v:    V{1, 1},
			u:    V{10, 11},
			buf:  V{999, 999},
			want: V{11, 12},
		},
		{
			name: "InPlace",
			v:    v,
			u:    V{10, 11},
			buf:  v,
			want: V{11, 12},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := Add(c.v, c.u, c.buf); got != c.want {
				t.Errorf("Add() = %v, want = %v", got, c.want)
			}
		})
	}
}
