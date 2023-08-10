package uploader

import (
	"context"
	"os"
	"strings"

	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/core"
	"github.com/dronestock/file/internal/internal"
	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/rickb777/gowebdav"
	"github.com/rickb777/gowebdav/auth"
)

var _ core.Uploader = (*Webdav)(nil)

type Webdav struct {
	*drone.Base
	gowebdav.Client

	addr     string
	username string
	password string
}

func NewWebdav(base *drone.Base, addr string, username string, password string) *Webdav {
	return &Webdav{
		Base:   base,
		Client: gowebdav.NewClient(addr, gowebdav.SetAuthentication(auth.Basic(username, password))),

		addr:     addr,
		username: username,
		password: password,
	}
}

func (w *Webdav) Mkdir(_ context.Context, dir string, permission os.FileMode) (err error) {
	if _, se := w.Stat(dir); nil != se {
		err = w.MkdirAll(dir, permission)
	}

	return
}

func (w *Webdav) Upload(ctx context.Context, path string, dir string, name string, _ os.FileMode) (err error) {
	_args := args.New().Build()
	_args.Option("user", gox.StringBuilder(w.username, internal.SeparatorUser, w.password).String())
	_args.Option("upload-file", path)
	_args.Add(strings.Join([]string{w.addr, dir, name}, internal.SeparatorPath))
	_, err = w.Command(internal.CommandCurl).Context(ctx).Args(_args.Build()).Build().Exec()

	return
}
