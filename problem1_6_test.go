package cracking_interview

import "testing"

func TestCompressString(t *testing.T) {
  tests := []struct {
    input string
    expected string
  }{
    {"", ""},
    {"a", "a1"},
    {"aa", "a2"},
    {"aaa", "a3"},
    {"aabcccccaaa", "a2b1c5a3"},
  }

  for _, test := range tests {
    result := CompressString(test.input)
    if result != test.expected {
      t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
    }
  }
}
