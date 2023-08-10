package uploader

import (
	"context"
	"fmt"
	"os"

	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/core"
	"github.com/dronestock/file/internal/internal"
	"github.com/goexl/gox/args"
)

var _ core.Uploader = (*Ssh)(nil)

type Ssh struct {
	*drone.Base

	addr     string
	username string
	password string
}

func NewSsh(base *drone.Base, addr string, username string, password string) *Ssh {
	return &Ssh{
		Base: base,

		addr:     addr,
		username: username,
		password: password,
	}
}

func (s *Ssh) Mkdir(_ context.Context, _ string, _ os.FileMode) (err error) {
	return
}

func (s *Ssh) Upload(ctx context.Context, path string, dir string, name string, _ os.FileMode) (err error) {
	_args := args.New().Build()
	_args.Arg("addr", s.addr)
	_args.Arg("username", s.username)
	_args.Arg("password", s.password)
	_args.Arg("local", path)
	_args.Arg("remote", fmt.Sprintf("%s/%s", dir, name))
	_, err = s.Command(internal.CommandScp).Context(ctx).Args(_args.Build()).Build().Exec()

	return
}
