package cracking_interview

// Note this is not the optimal implementation.
// The optimal implementation would only copy to the removeStack when necessary, and never copy back to the addStack.

type MyQueue struct {
  addStack Stack[int]
  removeStack Stack[int]
}

func (q *MyQueue) Add(item int) {
  for len(q.removeStack) > 0 {
    q.addStack.Push(q.removeStack.Pop())
  }
  q.addStack.Push(item)
}

func (q *MyQueue) Remove() int {
  for len(q.addStack) > 0 {
    q.removeStack.Push(q.addStack.Pop())
  }
  return q.removeStack.Pop()
}

func (q *MyQueue) Peek() int {
  for len(q.addStack) > 0 {
    q.removeStack.Push(q.addStack.Pop())
  }
  return q.removeStack.Peek()
}

func (q *MyQueue) Empty() {
  q.addStack = *new(Stack[int])
  q.removeStack = *new(Stack[int])
}
