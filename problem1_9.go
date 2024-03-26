package cracking_interview

import "strings"

func isRotation(s1, s2 string) bool {
  return strings.Contains(s2 + s2, s1)
}
