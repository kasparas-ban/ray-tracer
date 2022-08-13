package utils

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

type Color = Vec3
type Point3 = Vec3

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vec3) AddConst(t float64) Vec3 {
	return Vec3{v.X + t, v.Y + t, v.Z + t}
}

func (v Vec3) Mul(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v1 Vec3) MulVec(v2 Vec3) Vec3 {
	return Vec3{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

func (v Vec3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

// Util functions

func (v Vec3) ToString() string {
	return fmt.Sprintf("%v %v %v", v.X, v.Y, v.Z)
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v Vec3) Unit() Vec3 {
	unitLength := v.Length()
	return Vec3{
		v.X / unitLength,
		v.Y / unitLength,
		v.Z / unitLength,
	}
}

func RandomVec() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandomInInterval(min, max float64) Vec3 {
	return Vec3{RandomNum(min, max), RandomNum(min, max), RandomNum(min, max)}
}

func RandomInUnitSphere() Vec3 {
	var p Vec3
	for true {
		p = RandomInInterval(-1, 1)
		if p.LengthSq() >= 1 {
			continue
		}
		break
	}
	return p
}

func RandomUnitVec() Vec3 {
	return RandomInUnitSphere().Unit()
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Add(n.Mul(-v.Dot(n) * 2))
}

func Refract(uv, n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(uv.Mul(-1).Dot(n), 1.0)
	rOutPerp := uv.Add(n.Mul(cosTheta)).Mul(etaiOverEtat)
	rOutParallel := n.Mul(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSq())))
	return rOutPerp.Add(rOutParallel)
}

func RandomInUnitDisk() Vec3 {
	var p Vec3
	for true {
		p = Vec3{RandomNum(-1, 1), RandomNum(-1, 1), 0}
		if p.LengthSq() >= 1 {
			continue
		}
		break
	}
	return p
}
