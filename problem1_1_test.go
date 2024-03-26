package cracking_interview

import (
  "testing"
)

func TestHasUniqueChars(test *testing.T) {
  cases := []struct {
    input string
    expected bool
  }{
    {"", true},
    {"a", true},
    {"aa", false},
    {"ab", true},
    {"abc", true},
    {"abca", false},
  }
  for _, c := range cases {
    result := HasUniqueChars(c.input)
    if result != c.expected {
      test.Errorf("HasUniqueChars(%q) == %v, want %v", c.input, result, c.expected)
    }
  }
}
