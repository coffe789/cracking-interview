package cracking_interview

func ZeroMatrix(matrix [][]int) [][]int {
  zRows := make(map[int]bool)
  zCols := make(map[int]bool)

  for i := range(len(matrix)) {
    for j := range(len(matrix[0])) {
      if matrix[i][j] == 0 {
        zCols[j] = true
        zRows[i] = true
      }
    }
  }
  for i := range(len(matrix)) {
    for j := range(len(matrix[0])) {
      if zCols[j] || zRows[i] {
        matrix[i][j] = 0
      }
    }
  }
  return matrix
}


