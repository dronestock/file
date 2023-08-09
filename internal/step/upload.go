package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/config"
	"github.com/goexl/gox/field"
)

type Upload struct {
	*drone.Base

	servers []*config.Server
	uploads []*config.Upload
}

func NewUpload(base *drone.Base, servers []*config.Server, uploads []*config.Upload) *Upload {
	return &Upload{
		Base: base,

		servers: servers,
		uploads: uploads,
	}
}

func (u *Upload) Runnable() bool {
	return 0 != len(u.servers) && 0 != len(u.uploads)
}

func (u *Upload) Run(ctx context.Context) (err error) {
	last := len(u.servers) - 1
	for count, server := range u.servers {
		if ue := server.Upload(ctx, u.Base, u.uploads); nil != ue {
			err = ue
			u.Warn("上传文件出错", field.New("Server", server))
		}

		if nil != err && count != last {
			err = nil
		}
	}

	return
}
