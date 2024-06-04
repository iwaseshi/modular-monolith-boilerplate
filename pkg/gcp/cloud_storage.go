package gcp

import (
	"context"
	"io"
	"mime/multipart"
	"modular-monolith-boilerplate/pkg/adapter"
	"modular-monolith-boilerplate/pkg/restapi"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudStorage struct {
	Bucket *storage.BucketHandle
	client *storage.Client
	ctx    context.Context
	files  []file
}

type file struct {
	file        multipart.File
	filePath    string
	contentType string
}

func NewCloudStorage(ctx context.Context, bucket string) (adapter.Storage, error) {
	var client *storage.Client
	var err error
	if restapi.IsRunningOnCloud() {
		client, err = storage.NewClient(ctx)
	} else {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile("storage_key.json"))
	}
	if err != nil {
		return nil, err
	}
	_, err = client.Bucket(bucket).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return &CloudStorage{
		Bucket: client.Bucket(bucket),
		client: client,
		ctx:    ctx,
	}, nil
}

func (cs *CloudStorage) AddFile(addTarget multipart.File, filePath string, contentType string) {
	cs.files = append(cs.files, file{
		file:        addTarget,
		filePath:    filePath,
		contentType: contentType,
	})
}

func (st CloudStorage) Close() error {
	return st.client.Close()
}

func (cs *CloudStorage) WriteFiles() error {
	if len(cs.files) == 1 {
		return cs.writeFile(cs.files[0])
	}
	var wg sync.WaitGroup
	// エラーは1つだけ保存する
	errChan := make(chan error, 1)
	for _, f := range cs.files {
		wg.Add(1)
		go func(f file) {
			defer wg.Done()
			if err := cs.writeFile(f); err != nil {
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

func (cs CloudStorage) writeFile(f file) error {
	obj := cs.Bucket.Object(f.filePath)
	wc := obj.NewWriter(cs.ctx)
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
