package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"

	"github.com/jlaffaye/ftp"
)

type plugin struct {
	drone.Base
	*server

	// 服务器列表
	Servers []*server `default:"${SERVERS}"`
	// 上传文件
	Upload *upload `default:"${UPLOAD}"`
	// 上传文件列表
	Uploads []*upload `default:"${UPLOADS}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	if nil != p.Upload {
		if nil == p.Upload {
			p.Uploads = make([]*upload, 0, 1)
		}
		p.Uploads = append(p.Uploads, p.Upload)
	}

	// 初始化客户端
	if p.ftp, err = ftp.Dial(p.Addr, ftp.DialWithTimeout(p.Timeout)); nil == err {
		err = p.ftp.Login(p.Username, p.Password)
	}

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.upload, drone.Name("上传")),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("addr", p.Addr),
		field.New("username", p.Username),
		field.New("password", p.Password),

		field.New("uploads", p.Uploads),
	}
}
