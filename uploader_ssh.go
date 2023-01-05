package main

import (
	"fmt"
	"os"

	"github.com/dronestock/drone"
)

var _ uploader = (*uploaderSsh)(nil)

type uploaderSsh struct {
	addr     string
	username string
	password string
	plugin   *plugin
}

func newUploaderSsh(addr string, username string, password string, plugin *plugin) *uploaderSsh {
	return &uploaderSsh{
		addr:     addr,
		username: username,
		password: password,
		plugin:   plugin,
	}
}

func (us *uploaderSsh) mkdir(_ string, _ os.FileMode) (err error) {
	return
}

func (us *uploaderSsh) upload(path string, dir string, name string, _ os.FileMode) (err error) {
	args := []any{
		"--addr",
		us.addr,
		"--username",
		us.username,
		"--password",
		us.password,
		"--local",
		path,
		"--remote",
		fmt.Sprintf("%s/%s", dir, name),
	}
	err = us.plugin.Exec("scpx", drone.Args(args...))

	return
}
