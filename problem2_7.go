package cracking_interview

func doLlIntersect(ll1, ll2 *Node) bool {
  m1 := make(map[*Node]bool)
  m2 := make(map[*Node]bool)

  for ll1 != nil || ll2 != nil {
    // Check if seen
    if m1[ll2] {
      return true
    }
    if m2[ll1] {
      return true
    }

    // Store values
    m1[ll2] = true
    m2[ll1] = true

    // Iterate
    if ll1 != nil {
      ll1 = ll1.next
    }
    if ll2 != nil {
      ll2 = ll2.next
    }
  }

  // Final values were not checked
  if m1[ll2] {
    return true
  }
  if m2[ll1] {
    return true
  }

  return false
}
