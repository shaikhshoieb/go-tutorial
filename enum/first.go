package main

// import "fmt"

type serverState int

const (
	stateIdle serverState = iota
	stateConnected
	stateError
	stateRetrying
	stateUnkown
)

var stateName = map[serverState]string{
	stateIdle:      "Idle",
	stateConnected: "Connected",
	stateError:     "Error",
	stateRetrying:  "Retrying",
	stateUnkown:    "Unknown",
}

func (ss serverState) String() string {
	return stateName[ss]
}

func transition(s serverState) serverState {
	switch s {
	case stateIdle:
		return stateConnected
	case stateConnected, stateRetrying:
		return stateIdle
	case stateError:
		return stateError
	default:
		panic("Invalid state")
	}
}

func main() {
	ns := transition(stateIdle)
	println(ns.String())
	ns1 := transition(stateConnected)
	println(ns1.String())
	ns2 := transition(stateError)
	println(ns2.String())
	ns3 := transition(stateRetrying)
	println(ns3.String())
	// ns4 := transition(stateUnkown)
	// println(ns4.String())

}
