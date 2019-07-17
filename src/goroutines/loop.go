package main

var a string
var done bool

func setup() {
	a = "Hello, World"
	done = true
}

func main() {
	go setup() // kicks off a thread for setup()
	for !done {

	}
	print(a)
}

// on playground: Program exited: process took too long.
// in terminal:
// PS > go run loop.go
// Hello, World
