package cracking_interview

import (
  "testing"
)

func TestZeroMatrix(test *testing.T) {
  cases := []struct {
    input [][]int
    expected [][]int
  }{
    {
      [][]int{},
      [][]int{},
    },
    {
      [][]int{
        {1, 2, 3},
        {4, 0, 6},
        {7, 8, 9},
      },
      [][]int{
        {1, 0, 3},
        {0, 0, 0},
        {7, 0, 9},
      },
    },
    {
      [][]int{
        {1, 1, 1},
        {1, 1, 1},
        {1, 1, 1},
      },
      [][]int{
        {1, 1, 1},
        {1, 1, 1},
        {1, 1, 1},
      },
    },
  }
  for _, v := range(cases) {
    result := ZeroMatrix(v.input)
    for i := range(len(result)) {
      for j := range(len(result)) {
        if v.expected[i][j] != result[i][j] {
          test.Errorf("ZeroMatrix(%s) == %s, want %s", MatrixToString(v.input), MatrixToString(result), MatrixToString(v.expected))
        }
      }
    }
  }
}
