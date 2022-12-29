package main

type upload struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 目录
	Dir string `json:"dir"`
	// 文件名
	Filename string `json:"filename"`
	// 文件名列表
	Filenames []string `json:"filenames"`
}

func (p *plugin) upload() (undo bool, err error) {
	if undo = 0 == len(p.Uploads); undo {
		return
	}

	for _, _upload := range p.Uploads {
		if nil != _upload.Enabled && *_upload.Enabled {
			err = _upload.do(p)
		}

		if nil != err {
			return
		}
	}

	return
}
