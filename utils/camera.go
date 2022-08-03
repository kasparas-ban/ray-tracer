package utils

import "math"

type Camera struct {
	Origin          Point3
	Horizontal      Vec3
	Vertical        Vec3
	LowerLeftCorner Vec3
	U, V, W         Vec3
	lensRadius      float64
}

func (c Camera) GetRay(s, t float64) Ray {
	rd := RandomInUnitDisk().Mul(c.lensRadius)
	offset := c.U.Mul(rd.X).Add(c.V.Mul(rd.Y))

	return Ray{
		c.Origin.Add(offset),
		c.LowerLeftCorner.
			Add(c.Horizontal.Mul(s)).
			Add(c.Vertical.Mul(t)).
			Sub(c.Origin).
			Sub(offset),
	}
}

func GetCamera(
	lookfrom, lookat Point3,
	vup Vec3,
	vfov,
	aspectRatio, aperture, focusDist float64,
) Camera {
	theta := vfov * (math.Pi / 180)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookfrom.Sub(lookat).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)

	origin := lookfrom
	horizontal := u.Mul(viewportWidth).Mul(focusDist)
	vertical := v.Mul(viewportHeight).Mul(focusDist)
	lowerLeftCorner := origin.
		Sub(horizontal.Mul(0.5)).
		Sub(vertical.Mul(0.5)).
		Sub(w.Mul(focusDist))

	lensRadius := aperture / 2

	cam := Camera{
		origin,
		horizontal,
		vertical,
		lowerLeftCorner,
		u, v, w,
		lensRadius,
	}
	return cam
}
