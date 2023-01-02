package main

type uploader interface {
	mkdir(dir string) (err error)
	upload(filenames []string, dir string) (err error)
}
