package cracking_interview

// import "strings"
// func URLify(s string) string {
//   return strings.ReplaceAll(s, " ", "%20")
// }

func URLify(s string, l int) string {
  runes := []rune(s)
  space_count := 0

  for _, v := range(runes) {
    if v == ' ' {
      space_count += 1
    }
  }

  output_runes := make([]rune, l + (space_count * 2))
  output_idx := 0
  for _, v := range(runes) {
    if v != ' ' {
      output_runes[output_idx] = v
      output_idx += 1
    } else {
      output_runes[output_idx] = '%'
      output_runes[output_idx + 1] = '2'
      output_runes[output_idx + 2] = '0'
      output_idx += 3
    }
  }

  return string(output_runes)
}

// This does not work because you cannot resize an array of runes
// Note - I attempted this problem before I learned how to use slices
func URLifyInPlace(s []rune) []rune {
  space_count := 0
  for i := range(len(s)) {
    if s[i] == ' ' {
      space_count++
    }
  }

  for i := len(s) - 1; i >= 0; i-- {

  }

  return s
}
