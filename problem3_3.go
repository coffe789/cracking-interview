package cracking_interview

type StackSet struct {
  stacks Stack[Stack[int]]
  maxStackElements int
}

func StackSetNew(maxStackElements int) StackSet {
  var ret StackSet
  ret.stacks = append(ret.stacks, make(Stack[int], 0, maxStackElements))
  ret.maxStackElements = maxStackElements
  return ret
}

func (s *StackSet) activeStack() *Stack[int] {
  return &s.stacks[len(s.stacks) - 1]
}

func (s *StackSet) Push(element int) {
  // Increase stack count if needed
  if len(*s.activeStack()) == s.maxStackElements {
    s.stacks.Push(make(Stack[int], 0, s.maxStackElements))
  }

  // Add element to active stack
  (*s.activeStack()).Push(element)
}

func (s *StackSet) Pop() int {
  popVal := s.activeStack().Pop()

  // Reduce stack count if needed
  if len(*s.activeStack()) == 0 {
    s.stacks.Pop()
  }
  return popVal
}

// func (s *StackSet) PopAt(stackIndex int) int {
//   if stackIndex >= len(s.stacks) {
//     return 0
//   }
//   return s.stacks[stackIndex].Pop()
// }
