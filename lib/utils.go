package lib

import (
	"sync"
)

var counter int
var mu sync.Mutex

// getCounter will return the value of a global per-process incrementing counter
func getCounter() int {
	return count()
}

// resetCounter will reset incrementing counter's value
func resetCounter() {
	counter = 0
}

// count will increment the atomic counter and return its value
func count() int {
	mu.Lock()
	counter++
	mu.Unlock()
	return counter
}
