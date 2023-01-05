package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goexl/gox/field"
	"github.com/rickb777/gowebdav"
	"github.com/rickb777/gowebdav/auth"
)

var _ uploader = (*uploaderWebdav)(nil)

type uploaderWebdav struct {
	webdav gowebdav.Client
	plugin *plugin
}

func newUploaderWebdav(addr string, username string, password string, plugin *plugin) *uploaderWebdav {
	return &uploaderWebdav{
		webdav: gowebdav.NewClient(addr, gowebdav.SetAuthentication(auth.Basic(username, password))),
		plugin: plugin,
	}
}

func (uw *uploaderWebdav) mkdir(dir string, permission os.FileMode) (err error) {
	if _, se := uw.webdav.Stat(dir); nil != se {
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
			err = uw.webdav.WriteStream(fmt.Sprintf("%s/%s", dir, name), file, permission)
		}

		if nil != err {
			uw.plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
