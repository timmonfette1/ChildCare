package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  numChilds, _ := strconv.Atoi(os.Args[1])   // Number of children in the simulation
  numAdults, _ := strconv.Atoi(os.Args[2])   // Number of adults in the simulation

  // Channels for adults entering and leaving
  adultsEnter := make(chan int, numAdults)
  adultsLeave := make(chan int, numAdults)

  // Channels for children entering and leaving
  childEnter := make(chan int, numChilds)
  childLeave := make(chan int, numChilds)

  // Create children and adults
  fmt.Println("Starting Child Care Problem Simulation")
  go CreateChildren(childEnter, childLeave, numChilds)
  go CreateAdults(adultsEnter, adultsLeave, numAdults)

  // Start simulation
  StartDaycare(childEnter, adultsEnter, childLeave, adultsLeave)
}
