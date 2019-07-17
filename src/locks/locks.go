package main

import "sync"

var l sync.Mutex
var a string
var done bool

func setup() {
	a = "Hello, World"
	done = true
}

func main() {
	l.Lock()
	go setup()
	// l.Lock()
	// instructor says to use above lock, but if used both playground and terminal fail with error on that line.
	// playground err: fatal error: all goroutines are asleep - deadlock!
	// comment it out, both work as expected: playground times out, terminal prints message

	for !done {
	}

	print(a)
}
