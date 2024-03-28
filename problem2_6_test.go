package cracking_interview

import (
  "testing"
)

func TestIsLLPalindrome(t *testing.T) {
  head := &Node{value: 1}
  head.next = &Node{value: 2}
  head.next.next = &Node{value: 3}
  head.next.next.next = &Node{value: 2}
  head.next.next.next.next = &Node{value: 1}

  if !isLLPalindrome(head) {
    t.Errorf("Expected true, got false")
  }
}
