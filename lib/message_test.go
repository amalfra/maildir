package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

var mailDir string
var testData string
var msg *Message

func cleanMaildir() {
	err := os.RemoveAll(mailDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to clean maildir folder")
		os.Exit(1)
	}
}

func createMaildir() {
	for _, subDir := range Subdirs {
		err := os.MkdirAll(filepath.Join(mailDir, subDir), os.ModePerm)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create directory structure for maildir folder")
			os.Exit(1)
		}
	}
}

func init() {
	var err error
	mailDir = "/tmp/maildir-test"
	testData = "foo"
	cleanMaildir()
	createMaildir()

	msg, err = NewMessage(mailDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create message")
		os.Exit(1)
	}
}

func TestUnwrittenMessageDirShouldBeTemp(t *testing.T) {
	matched, _ := regexp.MatchString("(.*)tmp(.*)", msg.path())
	if msg.dir != "tmp" || !matched {
		t.Fatalf("unwritten message dir != tmp")
	}
}

func TestUnwrittenMessageHasUniquename(t *testing.T) {
	if len(msg.unqiueName) == 0 {
		t.Fatalf("unwritten message unqiueName empty")
	}
}

func TestUnwrittenMessageHasFilename(t *testing.T) {
	if len(msg.filename()) == 0 {
		t.Fatalf("unwritten message filename empty")
	}
}

func TestUnwrittenMessageHasNoInfo(t *testing.T) {
	if msg.info != "" {
		t.Fatalf("unwritten message info not empty")
	}
}

func TestUnwrittenMessageNotAbleToSetInfo(t *testing.T) {
	_, err := msg.SetInfo("test")
	if err == nil {
		t.Fatalf("unwritten message info shouldn't be updated")
	}
}

func TestCreateWrittenMessage(t *testing.T) {
	var err error
	cleanMaildir()
	createMaildir()

	msg, err = NewMessage(mailDir)
	if err != nil {
		t.Fatalf("failed to create message")
	}
	err = msg.Write(testData)
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to write to message. got error: %s", err))
	}
}

func TestWrittenMessageNotWritable(t *testing.T) {
	var err error
	err = msg.Write("noway!")
	if err == nil {
		t.Fatalf("it shouldn't be possbile to write to already written message")
	}
}

func TestWrittenMessageHaveNoInfo(t *testing.T) {
	if msg.info != "" {
		t.Fatalf("info should be empty")
	}
}

func TestWrittenMessageNotAbleToSetInfo(t *testing.T) {
	_, err := msg.SetInfo("test")
	if err == nil {
		t.Fatalf("written message info shouldn't be updated")
	}
}

func TestWrittenMessageDirShouldBeNew(t *testing.T) {
	matched, _ := regexp.MatchString("(.*)new(.*)", msg.path())
	if msg.dir != "new" || !matched {
		t.Fatalf("unwritten message dir != new")
	}
}

func TestWrittenMessageShouldHaveFile(t *testing.T) {
	if _, err := os.Stat(msg.path()); os.IsNotExist(err) {
		t.Fatalf("file should exist")
	}
}

func TestWrittenMessageHasCorrectData(t *testing.T) {
	data, err := msg.GetData()
	if err != nil {
		t.Fatalf("failed to read file")
	}
	if data != testData {
		t.Fatalf("incorrect data in file")
	}
}

func TestCreateProcessedMessage(t *testing.T) {
	var err error
	cleanMaildir()
	createMaildir()

	msg, err = NewMessage(mailDir)
	if err != nil {
		t.Fatalf("failed to create message")
	}
	err = msg.Write(testData)
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to write to message. got error: %s", err))
	}
	_, err = msg.Process()
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to process message. got error: %s", err))
	}
}

func TestProcessedMessageNotWritable(t *testing.T) {
	var err error
	err = msg.Write("noway!")
	if err == nil {
		t.Fatalf("it shouldn't be possbile to write to already processed message")
	}
}

func TestProcessedMessageDirShouldBeCur(t *testing.T) {
	matched, _ := regexp.MatchString("(.*)cur(.*)", msg.path())
	if msg.dir != "cur" || !matched {
		t.Fatalf("unwritten message dir != cur")
	}
}

func TestProcessedMessageHaveInfo(t *testing.T) {
	if msg.info != info {
		t.Fatalf("info not correct")
	}
}

func TestProcessedMessageAbleToSetInfo(t *testing.T) {
	_, err := msg.SetInfo("test-info")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to update info. got error: %s", err))
	}
	if msg.info != "test-info" {
		t.Fatalf("info not correct")
	}
	matched, _ := regexp.MatchString("(.*)test-info(.*)", msg.path())
	if !matched {
		t.Fatalf("correct info not present in message path")
	}
}

func TestProcessedMessageDestroy(t *testing.T) {
	err := msg.Destroy()
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to destroy message. got error: %s", err))
	}

	if _, err = os.Stat(msg.path()); err == nil {
		t.Fatalf(fmt.Sprintf("destroyed file still exists! got error: %s", err))
	}
}

func TestBadMessagePathErrorForData(t *testing.T) {
	cleanMaildir()
	createMaildir()

	_, err := msg.GetData()
	if err == nil {
		t.Fatalf("no error for bad message path")
	}
}

func TestBadMessagePathNotProcessed(t *testing.T) {
	_, err := msg.Process()
	if err == nil {
		t.Fatalf("no error when processing bad message path")
	}
}

func TestBadMessagePathResetMessageKey(t *testing.T) {
	oldKey := msg.Key()
	msg.Process()
	if oldKey != msg.Key() {
		t.Fatalf("message key not getting reset")
	}
}
