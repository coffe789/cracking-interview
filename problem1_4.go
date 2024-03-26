package cracking_interview

func isPermutationOfPalindrome(s string) bool {
  rune_counts := make(map[rune]int)
  
  for _, v := range(s) {
    if v != ' ' {
      rune_counts[v]++
    }
  }

  odd_counts := 0
  for _, v := range(rune_counts) {
    if (v % 2) != 0 {
      odd_counts++
    }
  }

  return odd_counts <= 1
}
