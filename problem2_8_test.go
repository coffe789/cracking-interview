package cracking_interview

import (
  "testing"
)

func TestDetectLoop(t *testing.T) {
  // Test case where there is no loop
  ll := &Node{value: 1, next: nil}
  ll.next = &Node{value: 1, next: nil}
  ll.next.next = &Node{value: 1, next: nil}
  ll.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next.next.next.next.next = &Node{value: 1, next: nil}
  ll.next.next.next.next.next.next.next.next.next = &Node{value: 1, next: nil}

  if detectLoop(ll) != nil {
    t.Errorf("Expected no loop, but found one")
  }

  // Test case where there is a loop
  ll.next.next.next.next.next.next.next.next.next = ll.next.next.next.next 

  if detectLoop(ll) != ll.next.next.next.next {
    t.Errorf("Incorrect loop detection")
  }
}
