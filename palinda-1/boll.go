package main

import (
    "fmt"
    "strconv"
    "time"
)

// TODO
// Implementera metoden så att den skickar "bollar"
// till kanalen.
// Funktionen ska skriva ut när den skickar en boll.
// Funktionen ska även skriva ut numret på bollen.
// Tänk på vilken typ bollen ska ha.
func kastare(ch chan<- string) {
    // ...
}
 
// TODO
// Implementera metoden så att den tar emot "bollar"
// från kanalen.
// Funktionen ska skriva ut när den tar emot en boll.
// Funktionen ska även skriva ut numret på bollen.
func fångare(ch <-chan string) {
    // ...
}

func main() {
    c := make(chan string)
    go kastare(c)
    go fångare(c)
    select {} // För att inte lämna main()
}

