package maildir

import (
	"os"
	"path"

	"github.com/amalfra/maildir/lib"
)

// Maildir implements maildir format and it's operations
type Maildir struct {
	path string
}

// NewMaildir will create new maildir at specified path
func NewMaildir(path string, create bool) *Maildir {
	maildir := new(Maildir)
	maildir.path = path
	if create {
		maildir.createDirectories()
	}
	return maildir
}

// createDirectories will the sub directories required by maildir
func (m *Maildir) createDirectories() {
	for i := 0; i < len(lib.Subdirs); i++ {
		os.MkdirAll(path.Join(m.path, lib.Subdirs[i]), os.ModePerm)
	}
}
