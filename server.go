package main

import (
	"time"

	"github.com/goexl/gox/field"
)

type server struct {
	// 类型
	Type serverType `default:"1" json:"type" validate:"required,oneof=1 2"`
	// 地址
	Addr string `json:"addr" validate:"required,hostname_port|url"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 超时
	Timeout time.Duration `default:"5s" json:"timeout"`
}

func (s *server) upload(plugin *plugin, uploads []*upload) (err error) {
	var _uploader uploader
	switch s.Type {
	case serverTypeWebdav:
		_uploader = newUploaderWebdav(s.Addr, s.Username, s.Password, plugin)
	case serverTypeFtp:
		_uploader, err = newUploaderFtp(s.Addr, s.Username, s.Password, s.Timeout, plugin)
	case serverTypeSsh:
		_uploader = newUploaderSsh(s.Addr, s.Username, s.Password, plugin)
	}
	if nil != err {
		return
	}

	// 遍历上传
	for _, _upload := range uploads {
		if nil == _upload.Enabled || !*_upload.Enabled {
			continue
		}

		if ue := _upload.do(_uploader, plugin); nil != ue {
			err = ue
			plugin.Warn("上传文件出错", field.New("upload", _upload), field.Error(ue))
		} else {
			plugin.Debug("上传文件成功", field.New("upload", _upload))
		}
	}

	return
}
