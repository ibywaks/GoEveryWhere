package GoEveryWhere

import (
	"fmt"
)

//Fire - fire an event
func Fire(eventName string, async bool) {
	listener, validEvent := Listen[eventName]

	if validEvent != true {
		fmt.Println("Event not registered!!")
		return
	}

	if async == true {
		go listener.Handle()
	} else {
		listener.Handle()
	}

}
