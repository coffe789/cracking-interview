package cracking_interview

import (
  "testing"
)

func TestMyQueue(t *testing.T) {
  q := MyQueue{}
  q.Add(1)
  q.Add(2)
  q.Add(3)
  q.Add(4)
  q.Add(5)

  if q.Remove() != 1 {
    t.Error("Expected 1")
  }
  if q.Remove() != 2 {
    t.Error("Expected 2")
  }
  if q.Remove() != 3 {
    t.Error("Expected 3")
  }
  if q.Remove() != 4 {
    t.Error("Expected 4")
  }
  if q.Remove() != 5 {
    t.Error("Expected 5")
  }
}
