package maildir

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

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

// Add writes data out as a new message. Returns Message instance
func (m *Maildir) Add(data string) (*lib.Message, error) {
	msg, err := lib.NewMessage(m.path)
	if err != nil {
		return nil, errors.New("failed to create message")
	}
	err = msg.Write(data)
	if err != nil {
		return nil, errors.New("failed to write message")
	}

	return msg, nil
}

// Get returns a message object for key
func (m *Maildir) Get(key string) *lib.Message {
	return lib.LoadMessage(m.path, key)
}

// List returns an arry of messages from new or cur directory, sorted by key
func (m *Maildir) List(dir string) (map[string]*lib.Message, error) {
	if !lib.StringInSlice(dir, lib.Subdirs) {
		return nil, errors.New("dir must be :new, :cur, or :tmp")
	}

	keys, err := m.getDirListing(dir)
	if err != nil {
		return nil, errors.New("failed to get directory listing")
	}
	sort.Sort(sort.StringSlice(keys))

	// map keys to message objects
	keyMap := make(map[string]*lib.Message)
	for i := 0; i < len(keys); i++ {
		keyMap[keys[i]] = m.Get(keys[i])
	}

	return keyMap, nil
}

// getDirListing returns an array of keys in dir
func (m *Maildir) getDirListing(dir string) ([]string, error) {
	filter := "*"
	searchPath := path.Join(m.path, dir, filter)
	filePaths, err := filepath.Glob(searchPath)
	// remove maildir path so that only key remains
	for i := 0; i < len(filePaths); i++ {
		filePaths[i] = strings.Replace(filePaths[i], m.path, "", 1)
	}
	return filePaths, err
}

// Delete a message by key
func (m *Maildir) Delete(key string) error {
	return m.Get(key).Destroy()
}
