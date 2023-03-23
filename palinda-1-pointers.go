package main

import "fmt"

func add(num1, num2 int) {
  num1 = num1 + num2
}

// TODO
// Gör en funktion addPointer som tar 2 argument: en int-pekare och en int.
// Funktionen ska addera int:en till pekarens värde.
// OBS! Funktionen ska INTE returnera någonting. :)

func main() {
  num := 5
  x := 10

  // Anropar add med x och num
  add(x, num)
  fmt.Printf("x efter add: %d\n\n", x)

  // TODO
  // Skapa en pekare som pekar på x

  // TODO
  // Anropa addPointer med din pekare och num

  fmt.Printf("x after addPointer: %d\n\n", x)
}
