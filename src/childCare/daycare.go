package main

import (
  "fmt"
)

var numA = 0
var numC = 0
var timeToExit = false

// Daycare Simulation
func StartDaycare(cEnter, aEnter, cLeave, aLeave chan int) {
  childEnter := make(chan int)
  adultLeave := make(chan int)

  for {
    // On-off "switch" for allowing children to enter the daycare
    if numC < numA * 3 {
      childEnter = cEnter
    } else {
      childEnter = nil
    }

    // On-off "switch" for allowing adults to leave the daycare
    if numC <= (numA - 1) * 3 {
      adultLeave = aLeave
    } else {
      adultLeave = nil
    }

    // Process all channels
    select {
    case id := <-aEnter:
      fmt.Printf("Adult %d entering the daycare\n", id)
      numA++
    case id := <-adultLeave:
      fmt.Printf("Adult %d leaving the daycare\n", id)
      numA--
      if (numA == 0) {      // When last adult leaves, it's time to exit
        timeToExit = true
      }
    case id := <-childEnter:
      fmt.Printf("Child %d entering the daycare\n", id)
      numC++
    case id := <-cLeave:
      fmt.Printf("Child %d leaving the daycare\n", id)
      numC--
    }

    // Check if it's time to exit the simulation
    if (timeToExit == true) {
      fmt.Println("All adults and children have gone home")
      fmt.Println("Ending child care simulation")
      break
    }
  }
}
