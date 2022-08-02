package utils

type Camera struct {
	AspectRatio    float64
	ViewportHeight float64
	ViewportWidth  float64
	FocalLength    float64

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
