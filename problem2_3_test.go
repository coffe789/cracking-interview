package cracking_interview

import (
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
  head := &Node{value: 1}
  head.next = &Node{value: 2}
  head.next.next = &Node{value: 3}
  head.next.next.next = &Node{value: 4}
  head.next.next.next.next = &Node{value: 5}

  DeleteMiddle(head.next.next)

  if head.next.next.value != 4 {
    t.Errorf("Expected 4, but got %v", head.next.next.value)
  }
}
