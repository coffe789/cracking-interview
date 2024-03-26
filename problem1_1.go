package cracking_interview

func HasUniqueChars(s string) bool {
  mymap := make(map[int32]bool)

  for _, c := range s {
    if !mymap[c] {
      mymap[c] = true
    } else {
      return false
    }
  }
  return true
}
