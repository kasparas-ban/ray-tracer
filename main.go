package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	. "example.com/utils"
	"github.com/schollz/progressbar/v3"
)

func randomScene() HittableList {
	var world HittableList

	groundMaterial := Lambertian{Color{0.5, 0.5, 0.5}}
	world.Append(Sphere{Point3{0, -1000, 0}, 1000, groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Point3{
				float64(a) + 0.9*rand.Float64(),
				0.2,
				float64(b) + 0.9*rand.Float64(),
			}

			if center.Sub(Point3{4, 0.2, 0}).Length() > 0.9 {
				var sphereMat Material

				if chooseMat < 0.8 {
					// Diffuse
					albedo := RandomVec().MulVec(RandomVec())
					sphereMat = Lambertian{albedo}
					world.Append(Sphere{center, 0.2, sphereMat})
				} else if chooseMat < 0.95 {
					// Metal
					albedo := Color{RandomNum(0.5, 1), RandomNum(0.5, 1), RandomNum(0.5, 1)}
					fuzz := RandomNum(0, 0.5)
					sphereMat = Metal{albedo, fuzz}
					world.Append(Sphere{center, 0.2, sphereMat})
				} else {
					// Glass
					sphereMat = Dielectric{1.5}
					world.Append(Sphere{center, 0.2, sphereMat})
				}
			}
		}
	}

	material1 := Dielectric{1.5}
	world.Append(Sphere{Point3{0, 1, 0}, 1.0, material1})

	material2 := Lambertian{Color{0.4, 0.2, 0.1}}
	world.Append(Sphere{Point3{-4, 1, 0}, 1.0, material2})

	material3 := Metal{Color{0.7, 0.6, 0.5}, 0.0}
	world.Append(Sphere{Point3{4, 1, 0}, 1.0, material3})

	return world
}

func rayColor(r Ray, world HittableList, depth int) Color {
	rec := HitRecord{}

	// If we've exceeded the ray bounce limit, no more light is gathered
	if depth <= 0 {
		return Color{0, 0, 0}
	}

	if didHit := world.Hit(r, 0.001, math.Inf(1), &rec); didHit {
		scattered := Ray{}
		attenuation := Color{}
		if rec.Mat.Scatter(r, rec, &attenuation, &scattered) {
			return attenuation.MulVec(rayColor(scattered, world, depth-1))
		}
		return Color{0, 0, 0}
	}
	unitDirection := r.Dir.Unit()
	t := 0.5 * (unitDirection.Y + 1)
	color1 := Color{1, 1, 1}
	color2 := Color{0.5, 0.7, 1}
	return color1.Mul(1 - t).Add(color2.Mul(t))
}

func main() {
	// Image

	aspectRatio := 3.0 / 2.0
	imageWidth := 1200
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samples := 100
	max_depth := 50

	// World

	world := randomScene()

	// Camera

	lookfrom := Point3{13, 2, 3}
	lookat := Point3{0, 0, 0}
	vup := Vec3{0, 1, 0}
	focusDist := 10.0
	aperture := 0.1

	cam := GetCamera(
		lookfrom,
		lookat,
		vup,
		20,
		aspectRatio,
		aperture,
		focusDist,
	)

	// Render
	f, _ := os.Create("image.ppm")
	defer f.Close()
	f.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight))
	bar := progressbar.Default(int64(imageHeight * imageWidth))

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			pixelColor := Color{0, 0, 0}
			for s := 0; s < samples; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)

				ray := cam.GetRay(u, v)
				pixelColor = pixelColor.Add(rayColor(ray, world, max_depth))
			}
			f.WriteString(WriteColor(pixelColor, samples))
			bar.Add(1)
		}
	}
}
