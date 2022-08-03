package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	. "example.com/utils"
	"github.com/schollz/progressbar/v3"
)

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

	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samples := 100
	max_depth := 50

	// World

	var world HittableList

	materialGround := Lambertian{Color{0.8, 0.8, 0.0}}
	materialCenter := Lambertian{Color{0.1, 0.2, 0.5}}
	materialLeft := Dielectric{1.5}
	materialRight := Metal{Color{0.8, 0.6, 0.2}, 0.0}

	world.Append(Sphere{Point3{0, -100.5, -1.0}, 100, materialGround})
	world.Append(Sphere{Point3{0, 0, -1.0}, 0.5, materialCenter})
	world.Append(Sphere{Point3{-1.0, 0, -1.0}, 0.5, materialLeft})
	world.Append(Sphere{Point3{-1.0, 0, -1.0}, -0.45, materialLeft})
	world.Append(Sphere{Point3{1.0, 0, -1.0}, 0.5, materialRight})

	// Camera

	lookfrom := Point3{3, 3, 2}
	lookat := Point3{0, 0, -1}
	vup := Vec3{0, 1, 0}
	focusDist := lookfrom.Sub(lookat).Length()
	aperture := 2.0

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
