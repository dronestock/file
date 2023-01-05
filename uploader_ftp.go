package main

import (
	"os"
	"time"

	"github.com/goexl/gfx"
	"github.com/jlaffaye/ftp"
)

var _ uploader = (*uploaderFtp)(nil)

type uploaderFtp struct {
	ftp *ftp.ServerConn
}

func newUploaderFtp(addr string, username string, password string, timeout time.Duration) (uf *uploaderFtp, err error) {
	uf = new(uploaderFtp)
	if uf.ftp, err = ftp.Dial(addr, ftp.DialWithTimeout(timeout)); nil == err {
		err = uf.ftp.Login(username, password)
	}

	return
}

func (uf *uploaderFtp) mkdir(_ string, _ os.FileMode) (err error) {
	return
}

func (uf *uploaderFtp) upload(path string, _ string, name string, _ os.FileMode) (err error) {
	if file, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = uf.ftp.Stor(gfx.Name(name), file)
	}

	return
}
