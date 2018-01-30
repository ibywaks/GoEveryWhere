package GoEveryWhere

import (
	"GoEveryWhere/listeners"
)

//Listen - map object registering events to listeners
var Listen = make(map[string]listeners.BaseListener)

//Register - function to register events to listeners
func Register(event string, listener listeners.BaseListener) {
	Listen[event] = listener
}
