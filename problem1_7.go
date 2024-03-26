package cracking_interview

func RotateMatrix(matrix [][]int) [][]int {
  var N = len(matrix)

  // Invert matrix
  var tmp int
  for i := range(N) {
    for j := range(N) {
      if i < j {
        tmp = matrix[i][j]
        matrix[i][j] = matrix[j][i]
        matrix[j][i] = tmp
      }
    }
  }

  // Flip matrix horizontally
  for i := range N {
    for j := range(N / 2) {
      flipIdx := N - 1 - j
      tmp = matrix[i][j]
      matrix[i][j] = matrix[i][flipIdx]
      matrix[i][flipIdx] = tmp
    }
  }

  return matrix
}
