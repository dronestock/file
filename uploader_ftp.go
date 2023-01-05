package main

import (
	"os"
	"time"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/jlaffaye/ftp"
)

var _ uploader = (*uploaderFtp)(nil)

type uploaderFtp struct {
	ftp    *ftp.ServerConn
	plugin *plugin
}

func newUploaderFtp(
	addr string,
	username string, password string,
	timeout time.Duration, plugin *plugin,
) (uf *uploaderFtp, err error) {
	uf = new(uploaderFtp)
	uf.plugin = plugin
	if uf.ftp, err = ftp.Dial(addr, ftp.DialWithTimeout(timeout)); nil == err {
		err = uf.ftp.Login(username, password)
	}

	return
}

func (uf *uploaderFtp) mkdir(dir string, permission os.FileMode) (err error) {
	return
}

func (uf *uploaderFtp) upload(filenames []string, dir string, permission os.FileMode) (err error) {
	for _, filename := range filenames {
		if file, oe := os.Open(filename); nil != oe {
			err = oe
		} else {
			err = uf.ftp.Stor(gfx.Name(filename), file)
		}

		if nil != err {
			uf.plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
