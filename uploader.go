package main

import (
	"os"
)

type uploader interface {
	mkdir(dir string, permission os.FileMode) (err error)
	upload(filenames []string, dir string, permission os.FileMode) (err error)
}
