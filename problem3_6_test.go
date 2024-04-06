package cracking_interview

import (
  "testing"
)

func TestAnimalQueue(t *testing.T) {
  EnqueueDog("dog1")
  EnqueueDog("dog2")
  EnqueueCat("cat1")
  EnqueueCat("cat2")
  EnqueueDog("dog3")
  EnqueueCat("cat3")
  EnqueueDog("dog4")
  EnqueueCat("cat4")
  EnqueueDog("dog5")
  EnqueueCat("cat5")
  EnqueueDog("dog6")
  EnqueueCat("cat6")
  
  ret := DequeueDog()
  if ret != "dog1" {
    t.Errorf("DequeueDog() = %v, want dog1", ret)
  }
  ret = DequeueCat()
  if ret != "cat1" {
    t.Errorf("DequeueCat() = %v, want cat1", ret)
  }
  ret = DequeueAny()
  if ret != "dog2" {
    t.Errorf("DequeueAny() = %v, want dog2", ret)
  }
  ret = DequeueAny()
  if ret != "cat2" {
    t.Errorf("DequeueAny() = %v, want cat2", ret)
  }
  ret = DequeueAny()
  if ret != "dog3" {
    t.Errorf("DequeueAny() = %v, want dog3", ret)
  }
}
