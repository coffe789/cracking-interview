package cracking_interview

func KthToLast(head *Node, k int) *Node {
  if head == nil {
    return nil
  }

  memory := make([]*Node, k+1, k+1)
  idx := 0

  for curr := head; curr != nil; curr = curr.next {
    memory[idx] = curr
    if curr.next == nil {
      kIdx := (idx + 1) % len(memory)
      return memory[kIdx]
    }
    idx = (idx + 1) % len(memory)
  }
  return nil
}
