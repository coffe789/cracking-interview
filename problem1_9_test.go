package cracking_interview

import (
  "testing"
)

func TestIsRotation(t *testing.T) {
  var cases = []struct {
    s1, s2 string
    expected bool
  }{
    {"", "", true},
    {"aaa", "aaa", true},
    {"waterbottle", "erbottlewat", true},
    {"waterbottle", "erbottlewaf", false},
  }
  for _, test := range cases {
    if output := isRotation(test.s1, test.s2); output != test.expected {
      t.Error("Test Failed: {} inputted, {} expected, received: {}", test.s1, test.s2, test.expected, output)
    }
  }
}
