package cracking_interview

import "strconv"

func MatrixToString(matrix [][]int) string {
  result := ""
  for i := range matrix {
    result += "["
    for j := range matrix[i] {
      result += strconv.Itoa(matrix[i][j])
      if j < len(matrix[i]) - 1 {
        result += " "
      }
    }
    result += "]"
  }
  return result
}

type Node struct {
  next *Node
  value int
}
