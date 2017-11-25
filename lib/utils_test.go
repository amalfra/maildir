package maildir

import (
	"testing"
	"time"
)

func TestGetCounter(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				getCounter()
			}
		}()
	}

	time.Sleep(time.Second)

	if getCounter() != 1000001 {
		t.Fatalf("Mutex not working as expected")
	}
}
