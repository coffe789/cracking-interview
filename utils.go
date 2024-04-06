package cracking_interview

import (
  "strconv"
)

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


// Stack //
type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
    *s = append(*s, value)
}

func (s *Stack[T]) Peek() T {
    if len(*s) == 0 {
        return *new(T) // Zero value
    }
    return (*s)[len(*s)-1]
}

func (s *Stack[T]) Pop() T {
    if len(*s) == 0 {
        return *new(T) // Zero value
    }
    value := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return value
}

func (s *Stack[T]) IsEmpty() bool {
  return len(*s) == 0
}
