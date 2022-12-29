package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// TODO 配置项，可以使用结构体
	// 如何配置选项请参数：https://github.com/dronestock/drone
	Todo string `default:"${PLUGIN_TODO=${TODO=默认值}}" validate:"required"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

// Steps TODO 返回所有要执行步骤
func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.todo, drone.Name(`启动守护进程`)),
	}
}

// Fields TODO 这儿返回所有的参数，上层在执行步骤时，会将参数在日志中打印
func (p *plugin) Fields() gox.Fields {
	return gox.Fields{
		field.String(`todo`, p.Todo),
	}
}
