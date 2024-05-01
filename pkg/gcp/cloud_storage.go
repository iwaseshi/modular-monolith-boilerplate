package gcp

import (
	"context"
	"io"
	"modular-monolith-boilerplate/pkg/adapter"
	"os"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudStorage struct {
	Bucket string
	client *storage.Client
	ctx    context.Context
	files  []file
}

type file struct {
	file        *os.File
	filePath    string
	contentType string
}

func NewCloudStorage(ctx context.Context, bucket string) (adapter.Storage, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("storage_key.json"))
	if err != nil {
		return nil, err
	}
	return &CloudStorage{
		Bucket: bucket,
		client: client,
		ctx:    ctx,
	}, nil
}

func (cs *CloudStorage) AddFile(addTarget *os.File, filePath string, contentType string) {
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
	obj := cs.client.Bucket(cs.Bucket).Object(f.filePath)
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
