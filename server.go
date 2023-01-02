package main

import (
	"time"
)

type server struct {
	// 类型
	Type serverType `default:"${TYPE=1}" json:"type" validate:"required,oneof=1 2"`
	// 地址
	Addr string `default:"${ADDR}" json:"addr" validate:"required,hostname_port"`
	// 用户名
	Username string `default:"${USERNAME}" json:"username" validate:"required"`
	// 密码
	Password string `default:"${Password}" json:"password" validate:"required"`
	// 超时
	Timeout time.Duration `default:"${TIMEOUT=5s}" json:"timeout"`
}

func (s *server) upload() (err error) {
	switch s.Type {
	case serverTypeWebdav:

	}

	return
}
