package main

import (
  "time"
  "math/rand"
)

// Function for creating adults every 3-5 seconds
func CreateAdults(enter, leave chan int, number int) {
  rand.Seed(time.Now().UnixNano())

  for i:= 0; i < number; i++ {
    x := rand.Intn(4) + 2
    time.Sleep(time.Second * time.Duration(x))
    go Child(enter, leave, i)
  }
}

// Adult goroutine (they want to leave after 20-30 seconds)
func Adult(enter, leave chan int, id int) {
  enter <- id
  x := rand.Intn(21) + 10
  time.Sleep(time.Second * time.Duration(x))
  leave <- id
}
