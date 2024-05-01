package adapter

import "os"

type Storage interface {
	AddFile(file *os.File, filename string, contentType string)
	WriteFiles() error
	Close() error
}
