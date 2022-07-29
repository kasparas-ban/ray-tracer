package utils

import (
  "fmt"
)

func WriteColor(c Color) string {
  return fmt.Sprintf("%v %v %v\n",
    uint(255 * c.X),
    uint(255 * c.Y),
    uint(255 * c.Z),
  )
}
