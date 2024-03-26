package cracking_interview

import "testing"

func TestURLify(test *testing.T) {
  cases := []struct {
    input string
    count int
    expected string
  }{
    {"abc", 3, "abc"},
    {"a b", 3, "a%20b"},
    {"a b c", 5, "a%20b%20c"},
    {"a b c d", 7, "a%20b%20c%20d"},
  }
  for _, c := range cases {
    result := URLify(c.input, c.count)
    if result != c.expected {
      test.Errorf("Expected %s, but got %s", c.expected, result)
    }
  }
}

