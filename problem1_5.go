package cracking_interview

func Abs(n int) int {
  if n < 0 {
    return -n
  }
  return n
}

func OneAway(s1, s2 string) bool {
  // Zero away
  if s1 == s2 {
    return true
  }
  // Not possible based on length
  if Abs(len(s1) - len(s2)) > 1 {
    return false
  }

  keys := make(map[rune]bool)
  rune_count1 := make(map[rune]int)
  rune_count2 := make(map[rune]int)

  for _, v := range(s1) {
    rune_count1[v]++
    keys[v] = true
  }

  for _, v := range(s2) {
    rune_count2[v]++
    keys[v] = true
  }

  hi_diff := 0
  lo_diff := 0
  for k := range(keys) {
    if rune_count1[k] > rune_count2[k] {
      hi_diff++
    } else if rune_count2[k] > rune_count1[k] {
      lo_diff++
    }
  }

  // Insert or remove char
  if (hi_diff == 1 && lo_diff == 0) || (hi_diff == 0 && lo_diff == 1) {
    return true
  }
  // Replace char
  if hi_diff == 1 && lo_diff == 1 {
    return true
  }
  return false
}
