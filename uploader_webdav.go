package main

import (
	"fmt"
	"os"

	"github.com/rickb777/gowebdav"
	"github.com/rickb777/gowebdav/auth"
)

var _ uploader = (*uploaderWebdav)(nil)

type uploaderWebdav struct {
	webdav gowebdav.Client
}

func newUploaderWebdav(addr string, username string, password string) *uploaderWebdav {
	return &uploaderWebdav{
		webdav: gowebdav.NewClient(addr, gowebdav.SetAuthentication(auth.Basic(username, password))),
	}
}

func (uw *uploaderWebdav) mkdir(dir string, permission os.FileMode) (err error) {
	if _, se := uw.webdav.Stat(dir); nil != se {
		err = uw.webdav.MkdirAll(dir, permission)
	}

	return
}

func (uw *uploaderWebdav) upload(path string, dir string, name string, permission os.FileMode) (err error) {
	if file, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = uw.webdav.WriteStream(fmt.Sprintf("%s/%s", dir, name), file, permission)
	}

	return
}
