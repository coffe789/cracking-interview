package cracking_interview

// import "fmt"

func partitionLL(lo, hi *Node) (*Node, *Node) {
  pivot := hi
  
  prev := lo
  for curr := lo.next; curr != nil ; curr = curr.next {
    if curr.value > pivot.value {
      // Move curr to end
      if curr == lo {
        lo = curr.next
      }

      prev.next = curr.next
      hi.next = curr
      curr.next = nil
      hi = curr
      curr = prev
    } else {
      prev = curr
    }

    if curr == pivot {
      return lo, curr
    }
  }
  return lo, pivot
}

