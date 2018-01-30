package main

import (
	"GoEveryWhere"
	"fmt"
)

//VisitEvent - a new visit event
type VisitEvent struct {
	Visitor string
}

//VisitListener - visit event listener
type VisitListener struct {
	Event VisitEvent
}

func main() {
	var visit = VisitEvent{"John Snow"}
	GoEveryWhere.Register("Visit Event", VisitListener{visit})

	GoEveryWhere.Fire("Visit Event", false)

	fmt.Println("Running some gnarly events!")
}

//Handle - Event handler method
func (listener VisitListener) Handle() {
	fmt.Println("Say hello to my little friend " + listener.Event.Visitor)
}

func (event VisitEvent) init() {
}
