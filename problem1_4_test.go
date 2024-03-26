package cracking_interview

import "testing"

func TestIsPermutationOfPalindrome(test *testing.T) {
  cases := []struct {
    input string
    expected bool
  }{
    {"", true},
    {"tact coa", true},
    {"aabb", true},
    {"123451234", true},
    {"1234512345", true},
    {"aaabbb", false},
    {"abbb", false},
  }

  for _, v := range(cases) {
    result := isPermutationOfPalindrome(v.input)
    if result != v.expected {
      test.Errorf("For input %s: expected %t, but got %t", v.input, v.expected, result)
    }
  }
}
