package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/goexl/gox/field"
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

func (us *uploaderSsh) upload(filenames []string, dir string, _ os.FileMode) (err error) {
	for _, filename := range filenames {
		_, name := filepath.Split(filename)
		args := []any{
			"--addr",
			us.addr,
			"--username",
			us.username,
			"--password",
			us.password,
			"--local",
			filename,
			"--remote",
			fmt.Sprintf("%s/%s", dir, name),
		}

		if err = us.plugin.Exec("scpx", drone.Args(args...));nil!= err {
			us.plugin.Warn("上传文件出错", field.New("filename", filename), field.Error(err))
		}
	}

	return
}
