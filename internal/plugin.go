package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/config"
	"github.com/dronestock/file/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Plugin struct {
	drone.Base
	config.Upload
	config.Server `validate:"omitempty,required_without=Servers"`

	// 服务器列表
	Servers []*config.Server `default:"${SERVERS}" validate:"omitempty,required_without=Server"`
	// 上传文件列表
	Uploads []*config.Upload `default:"${UPLOADS}"`
}

func NewPlugin() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Setup() (err error) {
	if nil == p.Servers {
		p.Servers = make([]*config.Server, 0, 1)
	}
	if p.Server.Set() {
		p.Servers = append(p.Servers, &p.Server)
	}

	if nil == p.Uploads {
		p.Uploads = make([]*config.Upload, 0, 1)
	}
	if p.Upload.Set() {
		p.Uploads = append(p.Uploads, &p.Upload)
	}

	return
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewUpload(&p.Base, p.Servers, p.Uploads)).Name("上传").Interrupt().Build(),
	}
}

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("servers", p.Servers),
		field.New("uploads", p.Uploads),
	}
}
