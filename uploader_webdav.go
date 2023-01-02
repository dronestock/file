package main

import (
	"os"
	"path/filepath"

	"github.com/goexl/gox/field"
	"github.com/rickb777/gowebdav"
	"github.com/rickb777/gowebdav/auth"
)

type uploaderWebdav struct {
	webdav gowebdav.Client
}

func newUploaderWebdav(addr string, username string, password string) *uploaderWebdav {
	return &uploaderWebdav{
		webdav: gowebdav.NewClient(addr, gowebdav.SetAuthentication(auth.Deferred(username, password))),
	}
}

func (uw *uploaderWebdav) mkdir(dir string, permission os.FileMode) (err error) {
	if _, se := uw.webdav.Stat(dir); nil != se {
		err = se
	} else if os.IsNotExist(se) {
		err = uw.webdav.MkdirAll(dir, permission)
	}

	return
}

func (uw *uploaderWebdav) upload(filenames []string, dir string, permission os.FileMode) (err error) {
	for _, filename := range filenames {
		if file, oe := os.Open(filename); nil != oe {
			err = oe
		} else {
			_, name := filepath.Split(filename)
			err = uw.webdav.WriteStream(filepath.Join(dir, name), file, permission)
		}

		if nil != err {
			plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
