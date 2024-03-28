package cracking_interview

func detectLoop(head *Node) *Node {
  seen := make(map[*Node]bool)

  for curr := head; curr != nil; curr = curr.next {
    if seen[curr] {
      return curr
    }
    seen[curr] = true
  }

  return nil
}
