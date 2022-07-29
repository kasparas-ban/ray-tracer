package utils

import (
  "fmt"
  "math"
)

type Vec3 struct {
  x, y, z float64
}

type Color = Vec3
type Point3 = Vec3

func NewVec3(x, y, z float64) Vec3 {
  return Vec3{x, y, z}
}

func AddVec3(v1, v2 Vec3) Vec3 {
  return Vec3{v1.x + v2.x, v1.y + v2.y, v1.z + v2.z}
}

func SubVec3(v1, v2 Vec3) Vec3 {
  return Vec3{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func MulVec3(v1, v2 Vec3) Vec3 {
  return Vec3{v1.x * v2.x, v1.y * v2.y, v1.z * v2.z}
}

func AddConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x + t, v.y + t, v.z + t}
}

func MulConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x * t, v.y * t, v.z * t}
}

func DivConstVec3(v Vec3, t float64) Vec3 {
  return Vec3{v.x / t, v.y / t, v.z / t}
}

func LengthSq(v Vec3) float64 {
  return math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2)
}

func Length(v Vec3) float64 {
  return math.Sqrt(LengthSq(v))
}

// Util functions

func PrintVec(v Vec3) string {
  return fmt.Sprintf("%v %v %v", v.x, v.y, v.z)
}

func Dot(v1, v2 Vec3) Vec3 {
  return Vec3{v1.x * v2.x, v1.y * v2.y, v1.z * v2.z}
}

func Cross(v1, v2 Vec3) Vec3 {
  return Vec3{
    v1.y * v2.z - v1.z * v2.y,
    v1.z * v2.x - v1.x * v2.z,
    v1.x * v2.y - v1.y * v2.x,
  }
}

func Unit(v Vec3) Vec3 {
  unit_length := Length(v)
  return Vec3{
    v.x / unit_length,
    v.y / unit_length,
    v.z / unit_length,
  }
}
