package cracking_interview

import (
  "testing"
)

func TestKthToLast(t *testing.T) {
  if KthToLast(nil, 0) != nil {
    t.Errorf("Test1: Expected nil, got %v", KthToLast(nil, 0))
  }
  
  ll := &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5, next: nil}}}}}
  if KthToLast(ll, 0).value != 5 {
    t.Errorf("Test2: Expected 5, got %d", KthToLast(ll, 0).value)
  }
  if KthToLast(ll, 1).value != 4 {
    t.Errorf("Test3: Expected 4, got %d", KthToLast(ll, 1).value)
  }
  if KthToLast(ll, 2).value != 3 {
    t.Errorf("Test4: Expected 3, got %d", KthToLast(ll, 2).value)
  }
  if KthToLast(ll, 3).value != 2 {
    t.Errorf("Test5: Expected 2, got %d", KthToLast(ll, 3).value)
  }
}
