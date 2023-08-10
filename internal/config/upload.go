package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal/core"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
)

type Upload struct {
	// 是否开启
	Enabled *bool `default:"${ENABLED=true}" json:"enabled,omitempty"`
	// 目录
	Dir string `default:"${DIR=.}" json:"dir,omitempty"`
	// 文件
	File string `default:"${FILE}" json:"filename,omitempty"`
	// 前缀
	Prefix string `default:"${PREFIX}" json:"prefix,omitempty"`
	// 后缀
	Suffix string `default:"${SUFFIX}" json:"suffix,omitempty"`
	// 文件权限
	Permission os.FileMode `default:"${PERMISSION=777}" json:"permission,omitempty"`
}

func (u *Upload) Set() bool {
	return "" != u.File
}

func (u *Upload) Do(ctx context.Context, uploader core.Uploader, base *drone.Base) (err error) {
	if err = uploader.Mkdir(ctx, u.Dir, u.Permission); nil != err {
		return
	}

	if paths, ae := gfx.All(".", gfx.Pattern(u.File)); nil != ae {
		err = ae
	} else if se := u.action(ctx, paths, uploader, base); nil != se {
		err = se
		base.Warn("上传文件出错", field.New("pattern", u.File), field.Error(err))
	} else {
		base.Debug("上传文件成功", field.New("pattern", u.File), field.Error(err))
	}

	return
}

func (u *Upload) action(ctx context.Context, paths []string, uploader core.Uploader, base *drone.Base) (err error) {
	for _, path := range paths {
		_, name := filepath.Split(path)
		ext := filepath.Ext(path)
		final := fmt.Sprintf("%s%s%s%s", u.Prefix, name[:len(name)-len(ext)], u.Suffix, ext)
		// 不在这里组装最终的路径而是把目录传到下一层，是因为怕各个上传服务器路径表现不一致
		// remote := fmt.Sprintf("%s/%s", u.Dir, final)

		if err = uploader.Upload(ctx, path, u.Dir, final, u.Permission); nil != err {
			base.Warn("上传文件出错", field.New("path", path), field.Error(err))
		}
	}

	return
}
