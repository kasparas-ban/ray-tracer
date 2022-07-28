package main

import(
  "fmt"
  "os"

  "example.com/utils"
  "github.com/schollz/progressbar/v3"
)

func main() {
  // Image
  image_width := 256
  image_height := 256

  // Render
  f, _ := os.Create("image.ppm")
  defer f.Close()
  f.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", image_width, image_height))
  bar := progressbar.Default(int64(image_height * image_width))

  for j := image_height-1; j >= 0; j-- {
    for i := 0; i < image_width; i++ {
      r := float64(i) /  float64(image_width - 1)
      g := float64(j) / float64(image_height - 1)
      b := 0.25

      ir := uint8(255 * r)
      ig := uint8(255 * g)
      ib := uint8(255 * b)

      line := fmt.Sprintf("%v %v %v\n", ir, ig, ib)
      f.WriteString(line)
      bar.Add(1)
    }
  }

  testVec := utils.NewVec3(10, 10, 10)
  fmt.Println(utils.Length(testVec))
}
