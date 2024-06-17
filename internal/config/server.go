package config

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/core"
	"github.com/dronestock/file/internal/uploader"
	"github.com/goexl/gox/field"
)

type Server struct {
	// 类型
	Type core.Type `default:"${TYPE=ssh}" json:"type,omitempty" validate:"required,oneof=ssh webdav ftp feishu lark"`
	// 基于安全连接
	SSH *Hosted `json:"ssh,omitempty" validate:"required_if=Type ssh"`
	// 基于文本协议
	Webdav *Hosted `json:"webdav,omitempty" validate:"required_if=Type webdav"`
	// 基于古老的文件服务
	Ftp *Hosted `json:"ftp,omitempty" validate:"required_if=Type ftp"`
	// 飞书
	Lark *Lark `json:"lark,omitempty" validate:"required_if=Type feishu | required_if=Type lark"`
}

func (s *Server) Set() bool {
	return "" != s.Addr
}

func (s *Server) Upload(ctx context.Context, base *drone.Base, uploads []*Upload) (err error) {
	var _uploader core.Uploader
	switch s.Type {
	case core.TypeWebdav:
		_uploader = uploader.NewWebdav(base, s.Addr, s.Username, s.Password)
	case core.TypeFtp:
		_uploader, err = uploader.NewFtp(s.Addr, s.Username, s.Password, s.Timeout)
	case core.TypeSsh:
		_uploader = uploader.NewSsh(base, s.Addr, s.Username, s.Password)
	}
	if nil != err {
		return
	}

	// 遍历上传
	for _, upload := range uploads {
		if nil == upload.Enabled || !*upload.Enabled {
			continue
		}

		if ue := upload.Do(ctx, _uploader, base); nil != ue {
			err = ue
			base.Warn("上传文件出错", field.New("upload", upload), field.Error(ue))
		} else {
			base.Debug("上传文件成功", field.New("upload", upload))
		}
	}

	return
}
