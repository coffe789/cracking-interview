package cracking_interview

func IsPermutation(s1, s2 string) bool {
  if len(s1) != len(s2) {
    return false
  }

  map1 := make(map[rune]int32)
  map2 := make(map[rune]int32)

  // Populate maps
  for i := range(len(s1)) {
    map1[rune(s1[i])] += 1
    map2[rune(s2[i])] += 1
  }

  // Test if maps are equal
  for key := range(map1) {
    if map1[key] != map2[key] {
      return false
    }
  }

  return true
}
