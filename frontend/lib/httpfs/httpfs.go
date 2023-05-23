package httpfs

import (
	"net/http"
	"strings"
)

type FileSystem struct {
	Fs http.FileSystem
}

func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.Fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.Fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
