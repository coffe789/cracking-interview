package cracking_interview

import "testing"

func TestIsPermutation(test *testing.T) {
  cases := []struct {
    s1, s2 string
    expected bool
  }{
    {"", "", true},
    {"a", "a", true},
    {"a", "b", false},
    {"ab", "ba", true},
    {"ab", "bc", false},
    {"abc", "abc", true},
    {"abc", "acb", true},
    {"abc", "bca", true},
    {"abc", "cab", true},
    {"abc", "cba", true},
    {"abc", "xyz", false},
  }
  for _, c := range cases {
    result := IsPermutation(c.s1, c.s2)
    if result != c.expected {
      test.Errorf("IsPermutation(%q, %q) == %v, want %v", c.s1, c.s2, result, c.expected)
    }
  }
}


