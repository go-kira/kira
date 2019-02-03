package session

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/go-kira/kog"
)

// FileHandler ...
type FileHandler struct {
	Path        string
	lifetime    int64
	filesSuffix string
	lock        sync.RWMutex
}

// NewFileHandler return FileHandler instance
func NewFileHandler(path string, lifetime int) *FileHandler {
	return &FileHandler{
		Path:     path,
		lifetime: int64(lifetime),
	}
}

// Read ...
func (f *FileHandler) Read(id string) ([]byte, error) {
	filename := filepath.Join(f.Path, "session_"+id)
	_, err := os.Stat(filename)
	if err != nil {
		// if there no file create new one, with empty data
		f.Write(id, nil)

		return nil, nil
	}

	// read the data from the file
	f.lock.RLock()
	defer f.lock.RUnlock()
	fdata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// return the raw data
	return fdata, nil
}

// Write ...
func (f *FileHandler) Write(id string, data []byte) error {
	// filename
	filename := filepath.Join(f.Path, "session_"+id)
	// lock
	f.lock.Lock()
	defer f.lock.Unlock()

	// write the file
	ioutil.WriteFile(filename, data, 0600)

	return nil
}

// Destroy ...
func (f *FileHandler) Destroy(id string) error {
	f.lock.RLock()
	defer f.lock.RUnlock()

	// file name
	filename := filepath.Join(f.Path, "session_"+id)

	// remove the session file
	return os.Remove(filename)
}

// GC to clean expired sessions.
func (f *FileHandler) GC() {
	f.lock.RLock()
	defer f.lock.RUnlock()

	// fetch all files inside the sessions folder, and check for expired files to delete them.
	if err := filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore hidding files.
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		if !info.IsDir() && (info.ModTime().Unix()+f.lifetime) < time.Now().Unix() {
			return os.Remove(path)
		}

		return nil
	}); err != nil {
		kog.Error("session garbage collecting")
	}
}
