package adapter

import "mime/multipart"

type Storage interface {
	AddFile(file multipart.File, filename string, contentType string)
	WriteFiles() error
	Close() error
}
