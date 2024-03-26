package cracking_interview

import (
	"testing"
)

func TestRotateMatrix(test *testing.T) {
  cases := []struct {
    input [][]int
    expected [][]int
  }{
    {
      [][]int{{1,2}, {3,4}},
      [][]int{{3,1}, {4,2}},
    },
    {
      [][]int{{1,2,3}, {4,5,6}, {7,8,9}},
      [][]int{{7,4,1}, {8,5,2}, {9,6,3}},
    },
    {
      [][]int{},
      [][]int{},
    },
  }

  for _, v := range(cases) {
    result := RotateMatrix(v.input)
    for i := range(len(result)) {
      for j := range(len(result)) {
        if v.expected[i][j] != result[i][j] {
          test.Errorf("RotateMatrix(%s) == %s, want %s", MatrixToString(v.input), MatrixToString(result), MatrixToString(v.expected))
        }
      }
    }
  }
}
