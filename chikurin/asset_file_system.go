package chikurin

import (
	"bytes"
	"net/http"
	"os"
	"strings"
)

type AssetFileSystem struct {
}

type AssetFile struct {
	*bytes.Reader
	os.FileInfo
}

func (fs AssetFileSystem) Open(name string) (http.File, error) {
	path := strings.TrimLeft(name, "/")

	data, err := Asset(path)
	if err != nil {
		return nil, err
	}

	info, _ := AssetInfo(path)
	file := &AssetFile{
		bytes.NewReader(data),
		info,
	}

	return file, nil
}

func (f *AssetFile) Close() error {
	return nil
}

func (f *AssetFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

func (f *AssetFile) Stat() (os.FileInfo, error) {
	return f, nil
}
