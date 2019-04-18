package crawler

import (
	"io"
)

type Crawler interface {
	Request(string) io.ReadCloser
	Parse(io.ReadCloser) interface{}
	Save(interface{}) error
}



