package cracking_interview

import "testing"

func TestSort(t *testing.T) {
  s := make(Stack[int], 0)
  s.Push(3)
  s.Push(1)
  s.Push(2)
  s.Push(5)
  s.Push(4)

  sort(&s)

  if s.Pop() != 1 {
    t.Errorf("Expected 1, got %d", s.Pop())
  }
  if s.Pop() != 2 {
    t.Errorf("Expected 2, got %d", s.Pop())
  }
  if s.Pop() != 3 {
    t.Errorf("Expected 3, got %d", s.Pop())
  }
  if s.Pop() != 4 {
    t.Errorf("Expected 4, got %d", s.Pop())
  }
  if s.Pop() != 5 {
    t.Errorf("Expected 5, got %d", s.Pop())
  }
}
