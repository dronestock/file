package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
)

type upload struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 目录
	Dir string `json:"dir"`
	// 文件名
	Filename string `json:"filename"`
	// 文件名列表
	Filenames []string `json:"filenames"`
	// 前缀
	Prefix string `json:"prefix"`
	// 后缀
	Suffix string `json:"suffix"`
	// 文件权限
	Permission os.FileMode `default:"0777" json:"permission"`
}

func (p *plugin) upload() (undo bool, err error) {
	if undo = 0 == len(p.Servers) || 0 == len(p.Uploads); undo {
		return
	}

	last := len(p.Servers) - 1
	for count, _server := range p.Servers {
		if ue := _server.upload(p, p.Uploads); nil != ue {
			err = ue
			p.Warn("上传文件出错", field.New("server", _server))
		}

		if nil != err && count != last {
			err = nil
		}
	}

	return
}

func (u *upload) do(uploader uploader, plugin *plugin) (err error) {
	if err = uploader.mkdir(u.Dir, u.Permission); nil != err {
		return
	}

	u.Filenames = append(u.Filenames, u.Filename)
	for _, filename := range u.Filenames {
		if names, ae := gfx.All(filename); nil != ae {
			err = ae
		} else if se := u.action(names, uploader, plugin); nil != se {
			err = se
			plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		} else {
			plugin.Debug("上传文件成功", field.New("filename", filename), field.Error(err))
		}
	}

	return
}

func (u *upload) action(filenames []string, uploader uploader, plugin *plugin) (err error) {
	for _, filename := range filenames {
		_, name := filepath.Split(filename)
		ext := filepath.Ext(filename)
		final := fmt.Sprintf("%s%s%s%s", u.Prefix, name[:len(name)-len(ext)], u.Suffix, ext)
		// 不在这里组装最终的路径而是把目录传到下一层，是因为怕各个上传服务器路径表现不一致

		if err = uploader.upload(filename, u.Dir, final, u.Permission); nil != err {
			plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
