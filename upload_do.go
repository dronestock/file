package main

import (
	"os"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
)

func (u *upload) do(plugin *plugin) (err error) {
	if err = plugin.ftp.ChangeDir(u.Dir); nil != err {
		return
	}

	u.Filenames = append(u.Filenames, u.Filename)
	for _, filename := range u.Filenames {
		if names, ae := gfx.All(filename); nil != ae {
			err = ae
		} else if se := u.store(plugin, names); nil != se {
			plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}

func (u *upload) store(plugin *plugin, filenames []string) (err error) {
	for _, filename := range filenames {
		if file, oe := os.Open(filename); nil != oe {
			err = oe
		} else {
			err = plugin.ftp.Stor(filename, file)
		}

		if nil != err {
			plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
