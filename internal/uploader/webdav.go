package uploader

import (
	"context"
	"fmt"
	"os"

	"github.com/dronestock/file/internal/core"
	"github.com/rickb777/gowebdav"
	"github.com/rickb777/gowebdav/auth"
)

var _ core.Uploader = (*Webdav)(nil)

type Webdav struct {
	webdav gowebdav.Client
}

func NewWebdav(addr string, username string, password string) *Webdav {
	return &Webdav{
		webdav: gowebdav.NewClient(addr, gowebdav.SetAuthentication(auth.Basic(username, password))),
	}
}

func (w *Webdav) Mkdir(_ context.Context, dir string, permission os.FileMode) (err error) {
	if _, se := w.webdav.Stat(dir); nil != se {
		err = w.webdav.MkdirAll(dir, permission)
	}

	return
}

func (w *Webdav) Upload(_ context.Context, path string, dir string, name string, permission os.FileMode) (err error) {
	if file, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = w.webdav.WriteStream(fmt.Sprintf("%s/%s", dir, name), file, permission)
	}

	return
}
