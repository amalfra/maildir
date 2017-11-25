package maildir

import (
	"sync"
)

var counter int
var mu sync.Mutex

// getCounter will return the value of a global per-process incrementing counter
func getCounter() int {
	return count()
}

// count will increment the atomic counter and return its value
func count() int {
	mu.Lock()
	counter++
	mu.Unlock()
	return counter
}
