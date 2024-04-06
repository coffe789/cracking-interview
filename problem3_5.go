package cracking_interview

func sort(s *Stack[int]) {
  if s.IsEmpty() {
    return
  }

  stackLen := len(*s)
  sortedCount := 0
  tmp := make(Stack[int], 0)

  for sortedCount < stackLen {
    hasHi := false
    hi := 0

    // Dump elements from s into tmp and store the maximum
    for !s.IsEmpty() {
      val := s.Pop()
      if val < hi || !hasHi {
        // If we are storing a lower maximum, put it back on the stack
        if hasHi {
          tmp.Push(hi)
        }
        hi = val
        hasHi = true
      } else {
        tmp.Push(val)
      }
    }

    // Put unsorted elements back on s
    for len(tmp) > sortedCount {
      s.Push(tmp.Pop())
    }

    // Put maximum on tmp
    tmp.Push(hi)
    sortedCount += 1
  }

  // Dump tmp back into s (this will invert it so we are sorted lowest to highest)
  for !tmp.IsEmpty() {
    s.Push(tmp.Pop())
  }
}
