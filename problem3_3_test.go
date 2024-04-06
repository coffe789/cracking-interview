package cracking_interview

import (
  "testing"
)

func TestStackSetNew(t *testing.T) {
  stackSet := StackSetNew(2)
  if len(stackSet.stacks) != 1 {
    t.Errorf("Expected 1 stack, got %d", len(stackSet.stacks))
  }
  if len(stackSet.stacks[0]) != 0 {
    t.Errorf("Expected 0 elements in stack, got %d", len(stackSet.stacks[0]))
  }

  stackSet.Push(1)
  stackSet.Push(1)
  stackSet.Push(1)
  if len(stackSet.stacks) != 2 {
    t.Errorf("Expected 2 stacks, got %d", len(stackSet.stacks))
  }
  if len(stackSet.stacks[0]) != 2 {
    t.Errorf("Expected 2 elements in stack, got %d", len(stackSet.stacks[0]))
  }
  if len(stackSet.stacks[1]) != 1 {
    t.Errorf("Expected 1 elements in stack, got %d", len(stackSet.stacks[1]))
  }

  stackSet.Pop()
  stackSet.Pop()
  if len(stackSet.stacks) != 1 {
    t.Errorf("Expected 1 stack, got %d", len(stackSet.stacks))
  }
  if len(stackSet.stacks[0]) != 1 {
    t.Errorf("Expected 2 elements in stack, got %d", len(stackSet.stacks[0]))
  }
}


