package cracking_interview

import "strings"

func isRotation(s1, s2 string) bool {
  if len(s1) != len(s2) {
    return false
  }
  return strings.Contains(s2 + s2, s1)
}
