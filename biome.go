package mapgen

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func getColor(transition bool, e float64, h1 string, c1e float64, h2 string, c2e float64) (colorful.Color, error) {
	c1, errorValue := colorful.Hex(h1)
	if transition {
		c2, _ := colorful.Hex(h2)
		return c1.BlendLab(c2, (e-c1e)/(c2e-c1e)).Clamped(), errorValue
	}

	return c1, errorValue
}

func biome(e float64, transition bool) (color.Color, error) {
	// #2c52a0; #3766c8; #d0d080; #589619; #426220; #5c453e; #4d3b39; #ffffff
	if e < 0.3 {
		return getColor(transition, e, "#0a1778", 0, "#3766c8", 0.3)
	} else if e < 0.4 {
		return getColor(transition, e, "#3766c8", 0.3, "#d0d080", 0.4)
	} else if e < 0.45 {
		return getColor(transition, e, "#d0d080", 0.4, "#589619", 0.45)
	} else if e < 0.60 {
		return getColor(transition, e, "#589619", 0.45, "#5c453e", 0.60)
	} else if e < 0.70 {
		return getColor(transition, e, "#5c453e", 0.60, "#4d3b39", 0.70)
	} else if e < 0.80 {
		return getColor(transition, e, "#4d3b39", 0.70, "#ffffff", 0.84)
	} else {
		return colorful.Hex("#ffffff")
	}
}
