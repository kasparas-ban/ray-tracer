package utils

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(Ray, HitRecord, *Color, *Ray) bool
}

type Lambertian struct {
	Albedo Color
}

type Metal struct {
	Albedo Color
	Fuzz   float64
}

type Dielectric struct {
	Ir float64
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDir := rec.Normal.Add(RandomUnitVec())

	// Catch degenerate scatter direction
	if scatterDir.NearZero() {
		scatterDir = rec.Normal
	}

	*scattered = Ray{rec.P, scatterDir}
	*attenuation = l.Albedo
	return true
}

func (l Metal) Scatter(rIn Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := Reflect(rIn.Dir.Unit(), rec.Normal)
	*scattered = Ray{rec.P, reflected.Add(RandomInUnitSphere().Mul(l.Fuzz))}
	*attenuation = l.Albedo
	return scattered.Dir.Dot(rec.Normal) > 0
}

func (d Dielectric) Scatter(rIn Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	*attenuation = Color{1.0, 1.0, 1.0}
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = 1.0 / d.Ir
	} else {
		refractionRatio = d.Ir
	}

	unitDir := rIn.Dir.Unit()
	cosTheta := math.Min(unitDir.Mul(-1).Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0

	var direction Vec3
	if cannotRefract || Reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = Reflect(unitDir, rec.Normal)
	} else {
		direction = Refract(unitDir, rec.Normal, refractionRatio)
	}

	*scattered = Ray{rec.P, direction}
	return true
}

func Reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reference
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
