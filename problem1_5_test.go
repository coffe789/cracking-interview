package cracking_interview

import "testing"

func TestOneAway(test *testing.T) {
  cases := []struct {
    s1, s2 string
    expected bool
  }{
    {"pale", "ple", true},
    {"pales", "pale", true},
    {"pale", "bale", true},
    {"pale", "bake", false},
  }

  for _, c := range cases {
    result := OneAway(c.s1, c.s2)
    if result != c.expected {
      test.Errorf("OneAway(%q, %q) == %t, expected %t", c.s1, c.s2, result, c.expected)
    }
  }
}
