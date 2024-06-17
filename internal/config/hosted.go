package config

import (
	"time"
)

type Hosted struct {
	// 地址
	Addr string `default:"${ADDR}" json:"addr,omitempty" validate:"required,ip|hostname|hostname_port|url"`
	// 用户名
	Username string `default:"${USERNAME}" json:"username,omitempty" validate:"required"`
	// 密码
	Password string `default:"${PASSWORD}" json:"password,omitempty" validate:"required"`
	// 超时
	Timeout time.Duration `default:"${TIMEOUT=5s}" json:"timeout,omitempty"`
}
