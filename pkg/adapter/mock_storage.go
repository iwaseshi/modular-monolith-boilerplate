package adapter

import (
	"context"
	"mime/multipart"
)

type MockStorage struct {
	NewBucketConnectionFunc func(ctx context.Context, bucket string) (Bucket, error)
}

func NewMockStorage() *MockStorage {
	mockBucket := &MockBucket{
		AddFileFunc: func(addTarget multipart.File, filePath string, contentType string) {
		},
		WriteFilesFunc: func() error {
			return nil
		},
		CloseFunc: func() error {
			return nil
		},
	}

	return &MockStorage{
		NewBucketConnectionFunc: func(ctx context.Context, bucket string) (Bucket, error) {
			return mockBucket, nil
		},
	}
}

func (m *MockStorage) NewBucketConnection(ctx context.Context, bucket string) (Bucket, error) {
	return m.NewBucketConnectionFunc(ctx, bucket)
}

type MockBucket struct {
	AddFileFunc    func(addTarget multipart.File, filePath string, contentType string)
	WriteFilesFunc func() error
	CloseFunc      func() error
}

func (mb *MockBucket) AddFile(addTarget multipart.File, filePath string, contentType string) {
	mb.AddFileFunc(addTarget, filePath, contentType)
}

func (mb *MockBucket) WriteFiles() error {
	return mb.WriteFilesFunc()
}

func (mb *MockBucket) Close() error {
	return mb.CloseFunc()
}
