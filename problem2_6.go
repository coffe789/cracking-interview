package cracking_interview

func isLLPalindrome(head *Node) bool {
  var l []int

  for curr := head; curr != nil; curr = curr.next {
    l = append(l, curr.value)
  }

  for i := 0; i < (len(l) / 2); i += 1 {
    if l[i] != l[len(l) - i - 1] {
      return false
    }
  }

  return true
}
