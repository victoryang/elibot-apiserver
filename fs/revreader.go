package fs

import (
	"io"
	"os"
)

type reverseReader struct {
	f *os.File
}

func (r reverseReader) Read(p []byte) (int, error) {
	l := len(p)
	_, err := r.f.Seek(int64(-l), os.SEEK_CUR)
	if err != nil {
		return 0, err
	}

	n, err := io.ReadFull(r.f, p)
	if err != nil {
		return n, err
	}

	_, err = r.f.Seek(int64(l), os.SEEK_CUR)
	return n, err
}

func newReverseReader(f *os.File) (reverseReader, error) {
	_, err := f.Seek(0, os.SEEK_END)
	return reverseReader{f}, err
}