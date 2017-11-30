package maildir

import (
	"fmt"
	"os"
	"path"
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

	maildir = NewMaildir(mailDir, true)
}

func TestPathProperty(t *testing.T) {
	if maildir.path != mailDir {
		t.Fatalf("path mismatch")
	}
}

func TestDirectoriesCreated(t *testing.T) {
	for i := 0; i < len(lib.Subdirs); i++ {
		if _, err := os.Stat(path.Join(mailDir, lib.Subdirs[i])); os.IsNotExist(err) {
			t.Fatalf("required sub directories not created")
		}
	}
}
