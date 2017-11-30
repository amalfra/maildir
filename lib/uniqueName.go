package maildir

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type uniqueName struct {
	now time.Time
}

// generate will create and return unique file names for new messages
func generate() (string, error) {
	uni := &uniqueName{now: time.Now()}
	hostname, err := uni.right()
	if err != nil {
		return "", errors.New("Failed to fetch hostname")
	}

	return fmt.Sprintf("%d.%s.%s", uni.left(), uni.middle(), hostname), nil
}

// left part of the unique name is the number of seconds since the UNIX epoch
func (uni *uniqueName) left() int64 {
	return uni.now.Unix()
}

func (uni *uniqueName) microseconds() int64 {
	return uni.now.Unix() * 1000000
}

// middle part of the unique name contains microsecond, process id, and a per-process incrementing counter
func (uni *uniqueName) middle() string {
	return fmt.Sprintf("M%06dP%dQ%d", uni.microseconds(), os.Getpid(), getCounter())
}

// right part is the hostname
func (uni *uniqueName) right() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return name, nil
}
