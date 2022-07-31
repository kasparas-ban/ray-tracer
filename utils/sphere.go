package utils

import (
	"math"
)

type Sphere struct {
	Center Point3
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64, rec HitRecord) (bool, HitRecord) {
	oc := r.Orig.Sub(s.Center)
	a := r.Dir.LengthSq()
	halfB := oc.Dot(r.Dir)
	c := oc.LengthSq() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false, rec
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false, rec
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).Mul(1 / s.Radius)
	rec = rec.setFaceNormal(r, outwardNormal)

	return true, rec
}
