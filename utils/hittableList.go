package utils

type Hittable interface {
	Hit(Ray, float64, float64, *HitRecord) bool
}

type HittableList []Hittable

func (hl *HittableList) clear() {
	hl = nil
}

func (hl *HittableList) Append(h Hittable) {
	*hl = append(*hl, h)
}

type HitRecord struct {
	p         Point3
	Normal    Vec3
	t         float64
	frontFace bool
}

func (h *HitRecord) setFaceNormal(r Ray, outwardNormal Vec3) {
	if frontFace := r.Dir.Dot(outwardNormal) < 0; frontFace {
		h.Normal = outwardNormal
	}
	h.Normal = outwardNormal.Mul(-1)
}

func (hl HittableList) Hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	var tempRec *HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, obj := range hl {
		if didHit := obj.Hit(r, tMin, closestSoFar, tempRec); didHit {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}

	return hitAnything
}
