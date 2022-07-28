package utils

import (
  "math"
)

type Vec3 struct {
  x float64
  y float64
  z float64
}

type Color Vec3
type Point3 Vec3

func NewVec3(x float64, y float64, z float64) Vec3 {
  return Vec3{x, y, z}
}

func addVec3(v1 Vec3, v2 Vec3) Vec3 {
  return Vec3{v1.x + v2.x, v1.y + v2.y, v1.z + v2.z}
}

func subVec3(v1 Vec3, v2 Vec3) Vec3 {
  return Vec3{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func mulVec3(v1 Vec3, v2 Vec3) Vec3 {
  return Vec3{v1.x * v2.x, v1.y * v2.y, v1.z * v2.z}
}

func addConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x + t, v.y + t, v.z + t}
}

func mulConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x * t, v.y * t, v.z * t}
}

func divConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x / t, v.y / t, v.z / t}
}

func lengthSq(v Vec3) float64 {
  return math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2)
}

func Length(v Vec3) float64 {
  return math.Sqrt(lengthSq(v))
}

