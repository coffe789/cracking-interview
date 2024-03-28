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

func partition(lo, hi *Node) (*Node, *Node, *Node) {
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
  }

  prev = lo
  for curr := lo.next; curr != nil ; curr = curr.next {
    if curr == pivot {
      return lo, prev, pivot
    }
  }
  return nil, nil, nil // This should never happen
}

func qs(lo, hi *Node) *Node {
  if lo == hi || lo == nil || hi == nil {
    return lo
  }

  lo, prePivot, pivot := partition(lo, hi)

  lo = qs(lo, prePivot)
  lo = qs(pivot.next, hi)

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
  for curr := ll.next; curr != nil; curr = curr.next {
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
