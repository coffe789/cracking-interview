package cracking_interview

type AnimalNode struct {
  nextOfType *AnimalNode
  next *AnimalNode
  prev *AnimalNode
  value string
  isDog bool
}

// Add to tail, remove from head
type AnimalQueue struct {
  head *AnimalNode
  tail *AnimalNode
}

var AQ = AnimalQueue{} // Animal
var DQ = AnimalQueue{} // Dog
var CQ = AnimalQueue{} // Cat

func EnqueueDog(name string) {
  tmp := &AnimalNode{value: name, isDog: true, prev: AQ.tail}

  if AQ.tail == nil {
    AQ.tail = tmp
    AQ.head = tmp
  } else {
    AQ.tail.next = tmp
    AQ.tail = tmp
  }
  if DQ.tail == nil {
    DQ.head = tmp
    DQ.tail = tmp
  } else {
    DQ.tail.nextOfType = tmp
    DQ.tail = tmp
  }
}

func EnqueueCat(name string) {
  tmp := &AnimalNode{value: name, isDog: false, prev: AQ.tail}

  if AQ.tail == nil {
    AQ.tail = tmp
    AQ.head = tmp
  }
  if CQ.tail == nil {
    CQ.head = tmp
    CQ.tail = tmp
  }

  AQ.tail.next = tmp
  AQ.tail = tmp

  CQ.tail.nextOfType = tmp
  CQ.tail = tmp
}

func DequeueDog() string {
  if DQ.head == nil {
    return ""
  }
  ret := DQ.head.value
  if DQ.head.prev == nil {
    AQ.head = DQ.head.next
    AQ.head.prev = nil
  } else {
    DQ.head.prev.next = DQ.head.next
  }
  DQ.head = DQ.head.nextOfType

  return ret
}

func DequeueCat() string {
  if CQ.head == nil {
    return ""
  }
  ret := CQ.head.value
  if CQ.head.prev == nil {
    AQ.head = CQ.head.next
    AQ.head.prev = nil
  } else {
    CQ.head.prev.next = CQ.head.next
  }
  CQ.head = CQ.head.nextOfType

  return ret
}

func DequeueAny() string {
  if AQ.head == nil {
    return ""
  }

  if AQ.head.isDog {
    return DequeueDog()
  }
  return DequeueCat()
}
