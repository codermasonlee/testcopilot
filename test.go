package main

import "fmt"

func print() {
	fmt.Println("xx")
}

type Run struct {
	Time    int // in milliseconds
	Results string
	Failed  bool
}

// Get average runtime of successful runs in seconds
func averageRuntimeInSeconds() {

}

func main() {
	print()
}
