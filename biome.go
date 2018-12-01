package mapgen

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type GradientTable []struct {
	Color      colorful.Color
	Pos        float64
	Transition bool
}

// This is the meat of the gradient computation. It returns a LAB-blend between
// the two colors around `t` respecting transition
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (self GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(self)-1; i++ {
		c1 := self[i]
		c2 := self[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			if c1.Transition {
				t := (t - c1.Pos) / (c2.Pos - c1.Pos)
				return c1.Color.BlendLab(c2.Color, t).Clamped()
			}

			if t-c1.Pos < c2.Pos-t {
				return c1.Color
			} else {
				return c2.Color
			}
		}
	}

	return self[len(self)-1].Color
}

// This is a very nice thing Golang forces you to do!
// It is necessary so that we can write out the literal of the colortable below.
func MustParseHex(s string) colorful.Color {
	c, _ := colorful.Hex(s)
	return c
}

func biome(e float64, transition bool) (color.Color, error) {

	keypoints := GradientTable{
		{MustParseHex("#0a1778"), 0.0, true},
		{MustParseHex("#3766c8"), 0.3, false},
		{MustParseHex("#d0d080"), 0.4, true},
		{MustParseHex("#589619"), 0.45, true},
		{MustParseHex("#5c453e"), 0.60, true},
		{MustParseHex("#4d3b39"), 0.70, true},
		{MustParseHex("#ffffff"), 0.80, true},
		{MustParseHex("#ffffff"), 1, true},
	}

	return keypoints.GetInterpolatedColorFor(e), nil
}
