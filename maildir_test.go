package maildir

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/amalfra/maildir/lib"
)

var mailDir string
var testData string
var maildir *Maildir

func cleanMaildir() {
	err := os.RemoveAll(mailDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to clean maildir folder")
		os.Exit(1)
	}
}

func init() {
	mailDir = "/tmp/maildir-test"
	testData = "foo"
	cleanMaildir()

	maildir = NewMaildir(mailDir)
}

func TestPathProperty(t *testing.T) {
	if maildir.path != mailDir {
		t.Fatalf("path mismatch")
	}
}

func TestDirectoriesCreated(t *testing.T) {
	for _, subDir := range lib.Subdirs {
		if _, err := os.Stat(filepath.Join(mailDir, subDir)); os.IsNotExist(err) {
			t.Fatalf("required sub directories not created")
		}
	}
}

func TestAdd(t *testing.T) {
	msg, err := maildir.Add(testData)
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to add message. got error: %s", err))
	}
	msgData, err := msg.GetData()
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get message data. got error: %s", err))
	}
	if msgData != testData {
		t.Fatalf("incorrect data saved in message")
	}
}

func TestListNewOneMessage(t *testing.T) {
	listing, err := maildir.List("new")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get messages listing. got error: %s", err))
	}

	for _, v := range listing {
		fileData, err := v.GetData()
		if err != nil {
			t.Fatalf(fmt.Sprintf("failed to read messages. got error: %s", err))
		}
		if fileData != testData {
			t.Fatalf("incorrect data saved in message")
		}
	}
}

func TestListNewMultipleMessage(t *testing.T) {
	TestAdd(t)
	listing, err := maildir.List("new")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get messages listing. got error: %s", err))
	}

	for _, v := range listing {
		fileData, err := v.GetData()
		if err != nil {
			t.Fatalf(fmt.Sprintf("failed to read messages. got error: %s", err))
		}
		if fileData != testData {
			t.Fatalf("incorrect data saved in message")
		}
	}
}

func TestProcessNewMessages(t *testing.T) {
	listing, err := maildir.List("cur")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get messages listing. got error: %s", err))
	}

	for _, v := range listing {
		_, err := v.Process()
		if err != nil {
			t.Fatalf(fmt.Sprintf("failed to process messages. got error: %s", err))
		}
	}
}

func TestListCurMultipleMessages(t *testing.T) {
	listing, err := maildir.List("cur")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get messages listing. got error: %s", err))
	}

	for _, v := range listing {
		fileData, err := v.GetData()
		if err != nil {
			t.Fatalf(fmt.Sprintf("failed to read messages. got error: %s", err))
		}
		if fileData != testData {
			t.Fatalf("incorrect data saved in message")
		}
	}
}

func TestDeleteMessages(t *testing.T) {
	listing, err := maildir.List("cur")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to get messages listing. got error: %s", err))
	}

	for _, v := range listing {
		err = maildir.Delete(v.Key())
		if err != nil {
			t.Fatalf(fmt.Sprintf("failed to delete messages. got error: %s", err))
		}
	}
}
