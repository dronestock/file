package main

import (
	"time"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"

	"github.com/jlaffaye/ftp"
)

type plugin struct {
	drone.Base

	// 地址
	Addr string `default:"${ADDR}" validate:"required,hostname_port"`
	// 用户名
	Username string `default:"${USERNAME}" validate:"required"`
	// 密码
	Password string `default:"${Password}" validate:"required"`
	// 超时
	Timeout time.Duration `default:"${TIMEOUT=5s}"`

	// 上传文件
	Upload *upload `default:"${UPLOAD}"`
	// 上传文件列表
	Uploads []*upload `default:"${UPLOADS}"`

	ftp *ftp.ServerConn
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
