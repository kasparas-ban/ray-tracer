package utils

import (
	"fmt"
)

func WriteColor(c Color, samples int) string {
	r := c.X
	g := c.Y
	b := c.Z

	// Divide the color by the number of samples
	scale := 1.0 / float64(samples)
	r *= scale
	g *= scale
	b *= scale

	return fmt.Sprintf("%v %v %v\n",
		int(256*Clamp(r, 0.0, 0.999)),
		int(256*Clamp(g, 0.0, 0.999)),
		int(256*Clamp(b, 0.0, 0.999)),
	)
}
