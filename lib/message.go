package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Message represents a maildir message and has it's supported operations
type Message struct {
	dir        string
	maildir    string
	info       string
	unqiueName string
	oldKey     string
}

// NewMessage will create new message in specified mail directory
func NewMessage(maildir string) (*Message, error) {
	var err error
	msg := new(Message)
	msg.maildir = maildir
	msg.dir = "tmp"
	msg.unqiueName, err = generate()
	if err != nil {
		return nil, errors.New("Failed to generate unqiue name")
	}
	return msg, nil
}

// parseKey will set dir, unqiueName, info based on the key
func (m *Message) parseKey(key string) {
	// remove leading /
	key = strings.TrimPrefix(key, string(os.PathSeparator))
	parts := strings.Split(key, string(os.PathSeparator))
	m.dir = parts[0]
	filename := parts[1]
	parts = strings.Split(filename, string(colon))
	m.unqiueName = parts[0]
	if len(parts) > 1 {
		m.info = parts[1]
	}
}

// LoadMessage will populate message object by loading info from passed key
func LoadMessage(maildir string, key string) *Message {
	msg := new(Message)
	msg.maildir = maildir
	msg.parseKey(key)
	return msg
}

// filename returns the filename of the message
func (m *Message) filename() string {
	return fmt.Sprintf("%s%c%s", m.unqiueName, colon, m.info)
}

// Key returns the key to identify the message
func (m *Message) Key() string {
	return filepath.Join(m.dir, m.filename())
}

// path returns the full path to the message
func (m *Message) path() string {
	return filepath.Join(m.maildir, m.Key())
}

// oldPath returns the old full path to the message
func (m *Message) oldPath() string {
	return filepath.Join(m.maildir, m.oldKey)
}

// rename the message. Returns the new key if successful
func (m *Message) rename(newDir string, newInfo string) (string, error) {
	// Save the old key so we can revert to the old state
	m.oldKey = m.Key()

	// Set the new state
	m.dir = newDir
	if newInfo != "" {
		m.info = newInfo
	}

	if m.oldPath() != m.path() {
		err := os.Rename(m.oldPath(), m.path())
		if err != nil {
			// restore old state
			if m.oldKey != "" {
				m.parseKey(m.oldKey)
			}
			return "", errors.New("Failed to rename folder")
		}
	}
	m.oldKey = ""
	return m.Key(), nil
}

// Write will write data to disk. only work with messages which haven't been written to disk.
// After successfully writing to disk, rename the message to new dir
func (m *Message) Write(data string) error {
	if m.dir != "tmp" {
		return errors.New("Can only write messages in tmp")
	}

	err := ioutil.WriteFile(m.path(), []byte(data), os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to write message to path %s", m.path())
	}

	_, err = m.rename("new", "")
	if err != nil {
		return errors.New("Failed to rename folder")
	}

	return nil
}

// Process will move a message from new to cur, add info. Returns the message's key
func (m *Message) Process() (string, error) {
	return m.rename("cur", info)
}

// SetInfo will set info on a message
func (m *Message) SetInfo(infoStr string) (string, error) {
	if m.dir != "cur" {
		return "", errors.New("Can only set info on cur messages")
	}
	return m.rename("cur", infoStr)
}

// GetData returns the message's data from disk
func (m *Message) GetData() (string, error) {
	dat, err := ioutil.ReadFile(m.path())
	strDat := string(dat)
	return strDat, err
}

// Destroy will remove the message file
func (m *Message) Destroy() error {
	return os.Remove(m.path())
}
