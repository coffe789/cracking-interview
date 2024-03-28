package cracking_interview

import "strconv"

func strIdx2Int(str string, idx int) int {
  if idx < 0 {
    return 0
  }
  return int(str[idx] - '0')
}

func sumLL(ll1, ll2 *Node) *Node {
  n1 := 0
  n2 := 0

  mult := 1
  for curr := ll1; curr != nil; curr = curr.next {
    n1 += curr.value * mult
    mult *= 10
  }

  mult = 1
  for curr := ll2; curr != nil; curr = curr.next {
    n2 += curr.value * mult
    mult *= 10
  }

  // Convert n3 to string
  n3_str := strconv.Itoa(n1 + n2)
  // Add last index to the head of ret
  ret := &Node{value: strIdx2Int(n3_str, len(n3_str) - 1), next: nil}

  // Add remaining indexes to tail
  curr := ret
  for i := len(n3_str) - 2; i >= 0; i -= 1 {
    curr.next = &Node{value: strIdx2Int(n3_str, i), next: nil}
    curr = curr.next
  }
  return ret
}
