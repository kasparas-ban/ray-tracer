package utils

type Hittable interface {
	Hit(Ray, float64, float64, HitRecord) (bool, HitRecord)
}

type HittableList []Hittable

func (hl *HittableList) clear() {
	hl = nil
}

func (hl *HittableList) Append(h Hittable) {
	*hl = append(*hl, h)
}

type HitRecord struct {
	P         Point3
	Normal    Vec3
	T         float64
	FrontFace bool
}

func (h HitRecord) setFaceNormal(r Ray, outwardNormal Vec3) HitRecord {
	if frontFace := r.Dir.Dot(outwardNormal) < 0; frontFace {
		h.Normal = outwardNormal
	}
	h.Normal = outwardNormal.Mul(-1)
	return h
}

func (hl HittableList) Hit(r Ray, tMin float64, tMax float64, rec HitRecord) (bool, HitRecord) {
	var tempRec HitRecord = HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, obj := range hl {
		if didHit, h := obj.Hit(r, tMin, closestSoFar, tempRec); didHit {
			hitAnything = true
			closestSoFar = h.T
			rec = h
		}
	}

	return hitAnything, rec
}
