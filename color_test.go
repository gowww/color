package color

import (
	"testing"
)

func TestColorHexAlpha(t *testing.T) {
	cases := []struct {
		color *Color
		want  string
	}{
		{&Color{0, 0, 0, 0}, "#0000"},
		{&Color{255, 255, 255, 1}, "#ffff"},
		{&Color{0, 75, 255, 1}, "#004bffff"},
		{&Color{0, 75, 255, .5}, "#004bff7f"},
	}
	for _, c := range cases {
		if got := c.color.HexAlpha(); c.want != got {
			t.Errorf("%v: want %v, got %v", c.color, c.want, got)
		}
	}
}
