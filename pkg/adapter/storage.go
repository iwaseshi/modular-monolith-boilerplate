package adapter

import (
	"context"
	"mime/multipart"
)

type Storage interface {
	NewBucketConnection(ctx context.Context, bucket string) (Bucket, error)
}

type Bucket interface {
	AddFile(file multipart.File, filename string, contentType string)
	WriteFiles() error
	Close() error
}
