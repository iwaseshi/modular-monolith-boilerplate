package gcp

import (
	"context"
	"io"
	"mime/multipart"
	"modular-monolith-boilerplate/pkg/adapter"
	"modular-monolith-boilerplate/pkg/di"
	"sync"

	"cloud.google.com/go/storage"
)

func init() {
	di.RegisterBean(NewCloudStorage)
}

type CloudStorage struct {
}

type bucket struct {
	handler *storage.BucketHandle
	client  *storage.Client
	ctx     context.Context
	files   []file
}

type file struct {
	file        multipart.File
	filePath    string
	contentType string
}

func NewCloudStorage() adapter.Storage {
	return &CloudStorage{}
}

func (cs *CloudStorage) NewBucketConnection(ctx context.Context, bucketName string) (adapter.Bucket, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	_, err = client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return &bucket{
		handler: client.Bucket(bucketName),
		client:  client,
		ctx:     ctx,
	}, nil
}

func (b *bucket) AddFile(addTarget multipart.File, filePath string, contentType string) {
	b.files = append(b.files, file{
		file:        addTarget,
		filePath:    filePath,
		contentType: contentType,
	})
}

func (b bucket) Close() error {
	return b.client.Close()
}

func (b *bucket) WriteFiles() error {
	if len(b.files) == 1 {
		return b.writeFile(b.files[0])
	}
	var wg sync.WaitGroup
	// エラーは1つだけ保存する
	errChan := make(chan error, 1)
	for _, f := range b.files {
		wg.Add(1)
		go func(f file) {
			defer wg.Done()
			if err := b.writeFile(f); err != nil {
				select {
				case errChan <- err:
				default:
				}
			}
		}(f)
	}
	wg.Wait()
	close(errChan)
	if err, ok := <-errChan; ok {
		return err
	}
	return nil
}

func (b bucket) writeFile(f file) error {
	obj := b.handler.Object(f.filePath)
	wc := obj.NewWriter(b.ctx)
	defer wc.Close()
	wc.ContentType = f.contentType
	if _, err := io.Copy(wc, f.file); err != nil {
		f.file.Close()
		return err
	}
	if err := f.file.Close(); err != nil {
		return err
	}
	return nil
}
