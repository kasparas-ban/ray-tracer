package utils

import "math"

type Camera struct {
	AspectRatio    float64
	ViewportHeight float64
	ViewportWidth  float64

	Origin          Point3
	Horizontal      Vec3
	Vertical        Vec3
	LowerLeftCorner Vec3
}

func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		c.Origin,
		c.LowerLeftCorner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)).Sub(c.Origin),
	}
}

func GetCamera(lookfrom, lookat Point3, vup Vec3, vfov, aspectRatio float64) Camera {
	theta := vfov * (math.Pi / 180)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookfrom.Sub(lookat).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)

	// focalLength := 1.0

	origin := lookfrom
	horizontal := u.Mul(viewportWidth)
	vertical := v.Mul(viewportHeight)
	lowerLeftCorner := origin.
		Sub(horizontal.Mul(0.5)).
		Sub(vertical.Mul(0.5)).
		Sub(w)

	cam := Camera{
		aspectRatio,
		viewportHeight,
		viewportWidth,
		origin,
		horizontal,
		vertical,
		lowerLeftCorner,
	}
	return cam
}
