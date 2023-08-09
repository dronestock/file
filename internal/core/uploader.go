package core

import (
	"context"
	"os"
)

type Uploader interface {
	Mkdir(ctx context.Context, dir string, permission os.FileMode) (err error)

	Upload(ctx context.Context, path string, dir string, name string, permission os.FileMode) (err error)
}
