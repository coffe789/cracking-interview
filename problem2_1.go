package cracking_interview

func DeleteDuplicates(ll Node) Node {
  seen := make(map[int]bool)

  previous := &ll
  for curr := ll.next; curr != nil; curr = curr.next {
    if seen[curr.value] {
      // Erase curr
      previous.next = curr.next
      curr = previous
    } else {
      seen[ll.value] = true
      previous = curr
    }
  }

  return ll
}

func partition(lo, hi *Node) (*Node, *Node) {
  pivot := hi
  
  for curr := lo; ; curr = curr.next {
    if curr.value > pivot.value {
      // Move curr to end
      tmp := curr
      curr = curr.next
      hi.next = tmp
      hi = tmp
    }
    if curr.next == pivot {
      return curr, pivot // Also return Node before pivot to use as next hi
    }
  }
}

func qs(lo, hi *Node) *Node {
  if lo == hi || lo == nil {
    return lo
  }

  prePivot, pivot := partition(lo, hi)
  qs(lo, prePivot)
  qs(pivot.next, hi)

  return lo
}

func DeleteDuplicates2(ll *Node) *Node {
  // Sort ll
  for curr := ll; curr != nil; curr = curr.next {
    if curr.next == nil {
      ll = qs(ll, curr)
    }
  }

  previous := ll
  for curr := ll.next; curr != nil ; curr = curr.next {
    if curr.value == curr.next.value {
      // Erase curr
      previous.next = curr.next
      curr = curr.next
    } else {
      previous = curr
    }
  }
  return ll
}
