package cracking_interview

import "testing"

func TestSumLL(t *testing.T) {
  ll1 := &Node{value: 7, next: &Node{value: 1, next: &Node{value: 6, next: nil}}}
  ll2 := &Node{value: 5, next: &Node{value: 9, next: &Node{value: 2, next: nil}}}
  ll3 := sumLL(ll1, ll2)
  if ll3.value != 2 || ll3.next.value != 1 || ll3.next.next.value != 9 {
    t.Errorf("Test1: Expected 2 -> 1 -> 9, got %d -> %d -> %d", ll3.value, ll3.next.value, ll3.next.next.value)
  }
}
