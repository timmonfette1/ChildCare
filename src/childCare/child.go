package main

import (
  "time"
  "math/rand"
)

// Function for creating children every 1-3 seconds
func CreateChildren(enter, leave chan int, number int) {
  rand.Seed(time.Now().UnixNano())

  for i:= 0; i < number; i++ {
    x := rand.Intn(2) + 1
    time.Sleep(time.Second * time.Duration(x))
    go Child(enter, leave, i)
  }
}

// Child goroutine (they want to leave after 10-15 seconds)
func Child(enter, leave chan int, id int) {
  enter <- id
  x := rand.Intn(11) + 5
  time.Sleep(time.Second * time.Duration(x))
  leave <- id
}
