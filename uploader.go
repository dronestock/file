package main

import (
	"os"
)

type uploader interface {
	mkdir(dir string, permission os.FileMode) (err error)
	upload(path string, dir string, name string, permission os.FileMode) (err error)
}
