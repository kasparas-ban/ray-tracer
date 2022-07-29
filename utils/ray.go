package utils

type Ray struct {
  Orig Point3
  Dir Vec3
}

func (r Ray) at(t float64) Point3 {
  return r.Orig.Add(r.Dir.AddConst(t))
}
