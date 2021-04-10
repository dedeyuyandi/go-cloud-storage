package storage

import (
	"context"
	"io"
)

type StorageUploadFile interface {
	io.Closer
	UploadFile(ctx context.Context, req interface{}) (string, error)
}
