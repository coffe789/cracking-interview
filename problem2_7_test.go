package cracking_interview

import "testing"

func TestDoLlIntersect(t *testing.T) {
  a := &Node{value: 1, next:nil}
  b := &Node{value: 2, next:a}
  c := &Node{value: 3, next:a}
  d := &Node{value: 4, next:b}

  if !doLlIntersect(d, c) {
    t.Error("Expected true, got false")
  }
}
