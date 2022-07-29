package utils

import (
  "fmt"
)

func WriteColor(c Color) string {
  return fmt.Sprintf("%v %v %v\n",
    uint(255 * c.x),
    uint(255 * c.y),
    uint(255 * c.z),
  )
}
