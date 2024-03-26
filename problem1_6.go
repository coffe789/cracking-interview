package cracking_interview

import "strconv"

func CompressString(s string) string {
  if len(s) == 0 {
    return s
  }

  output := string(s[0])
  count := 1

  for i := 1; i < len(s); i++ {
    if output[len(output) - 1] == s[i] {
      count++
    } else {
      output += strconv.Itoa(count)
      output += string(s[i])
      count = 1
    }
  }
  output += strconv.Itoa(count)

  return output
}
