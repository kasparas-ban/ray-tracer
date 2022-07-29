package main

import(
  "fmt"
  "os"

  . "example.com/utils"
  "github.com/schollz/progressbar/v3"
)

func rayColor(r Ray) Color {
  unitDirection := r.Dir.Unit()
  t := 0.5 * (unitDirection.Y + 1)
  color1 := Color { 1, 1, 1 }
  color2 := Color { 0.5, 0.7, 1 }
  return color1.Mul(1 - t).Add(color2.Mul(t))
}

func main() {
  // Image

  aspectRatio := 16.0 / 9.0
  imageWidth := 400
  imageHeight := int(float64(imageWidth) / aspectRatio)

  // Camera

  viewportHeight := 2.0
  viewportWidth := aspectRatio * viewportHeight
  focalLength := 1.0

  origin := Point3 { 0, 0, 0 }
  horizontal := Vec3 { viewportWidth, 0, 0 }
  vertical := Vec3 { 0, viewportHeight, 0 }
  lowerLeftCorner := origin.Sub(horizontal.Mul(1/2)).Sub(vertical.Mul(1/2)).Sub(Vec3{ 0, 0, focalLength })

  // Render
  f, _ := os.Create("image.ppm")
  defer f.Close()
  f.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight))
  bar := progressbar.Default(int64(imageHeight * imageWidth))

  for j := imageHeight-1; j >= 0; j-- {
    for i := 0; i < imageWidth; i++ {

      u := float64(i) / float64(imageWidth - 1)
      v := float64(j) / float64(imageHeight - 1)

      ray := Ray {
        origin,
        lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)).Sub(origin),
      }

      pixelColor := rayColor(ray)

      //pixelColor := Color {
      //  float64(i) /  float64(imageWidth - 1),
      //  float64(j) / float64(imageHeight - 1),
      //  0.25,
      //}

      f.WriteString(WriteColor(pixelColor))
      bar.Add(1)
    }
  }
}
