package cracking_interview

import "testing"

func TestDeleteDuplicates(t *testing.T) {
  ll := Node{value: 1, next: &Node{value: 2, next: &Node{value: 1, next: &Node{value: 3, next: &Node{value: 2, next: nil}}}}}
  ll = DeleteDuplicates(ll)
  if ll.value != 1 || ll.next.value != 2 || ll.next.next.value != 3 {
    t.Errorf("Test1: Expected 1 -> 2 -> 3, got %d -> %d -> %d", ll.value, ll.next.value, ll.next.next.value)
  }


  ll2 := &Node{value: 1, next: &Node{value: 2, next: &Node{value: 1, next: &Node{value: 3, next: &Node{value: 2, next: nil}}}}}
  ll2 = DeleteDuplicates2(ll2)
  if ll2.value != 1 || ll2.next.value != 2 || ll2.next.next.value != 3 {
    t.Errorf("Test2: Expected 1 -> 2 -> 3, got %d -> %d -> %d", ll2.value, ll2.next.value, ll2.next.next.value)
  }
}
