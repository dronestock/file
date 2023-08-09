package uploader

import (
	"context"
	"os"
	"time"

	"github.com/dronestock/file/internal/core"
	"github.com/goexl/gfx"
	"github.com/jlaffaye/ftp"
)

var _ core.Uploader = (*Ftp)(nil)

type Ftp struct {
	ftp *ftp.ServerConn
}

func NewFtp(addr string, username string, password string, timeout time.Duration) (uf *Ftp, err error) {
	uf = new(Ftp)
	if uf.ftp, err = ftp.Dial(addr, ftp.DialWithTimeout(timeout)); nil == err {
		err = uf.ftp.Login(username, password)
	}

	return
}

func (f *Ftp) Mkdir(_ context.Context, _ string, _ os.FileMode) (err error) {
	return
}

func (f *Ftp) Upload(_ context.Context, path string, _ string, name string, _ os.FileMode) (err error) {
	if file, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = f.ftp.Stor(gfx.Name(name), file)
	}

	return
}
