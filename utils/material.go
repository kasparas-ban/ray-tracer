package utils

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
