package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 服务器
	Server *server `default:"${SERVER}" validate:"omitempty,required_without=Servers"`
	// 服务器列表
	Servers []*server `default:"${SERVERS}" validate:"omitempty,required_without=Server"`
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
	if nil == p.Servers {
		p.Servers = make([]*server, 0, 1)
	}
	p.Servers = append(p.Servers, p.Server)

	if nil == p.Uploads {
		p.Uploads = make([]*upload, 0, 1)
	}
	if nil != p.Upload {
		p.Uploads = append(p.Uploads, p.Upload)
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
		field.New("servers", p.Servers),
		field.New("uploads", p.Uploads),
	}
}
