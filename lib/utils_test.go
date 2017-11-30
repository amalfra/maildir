package maildir

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCounter(t *testing.T) {
	resetCounter()
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				getCounter()
			}
		}()
	}

	time.Sleep(time.Second)
	countVal := getCounter()

	if countVal != 1000001 {
		t.Fatalf(fmt.Sprintf("Mutex not working as expected, counter returned %d", countVal))
	}
}
