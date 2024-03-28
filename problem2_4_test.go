package cracking_interview

import (
  "testing"
  "fmt"
)

func TestPartitionLL(t *testing.T) {
  ll := &Node{nil, 3}
  ll.next = &Node{ll, 5}
  ll.next.next = &Node{ll, 8}
  ll.next.next.next = &Node{ll, 5}
  ll.next.next.next.next = &Node{ll, 10}
  ll.next.next.next.next.next = &Node{ll, 2}
  ll.next.next.next.next.next.next = &Node{ll, 6}

  ll, _ = partitionLL(ll, ll.next.next.next.next.next.next)

  fmt.Println("Manual test with 6 as pivot:")
  for curr := ll; curr != nil; curr = curr.next {
    fmt.Print(curr.value, ",")
  }
  fmt.Println()
}
