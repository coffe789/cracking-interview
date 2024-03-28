package cracking_interview

func DeleteMiddle(middleNode *Node) {
  if middleNode == nil {
    return
  }
  *middleNode = *middleNode.next
}
