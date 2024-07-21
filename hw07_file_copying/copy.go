package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type ProgressReader struct {
	io.Reader
	bar *pb.ProgressBar
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	if n > 0 {
		for i := 0; i < n; i++ {
			pr.bar.Increment()
			// time.Sleep(time.Millisecond)
		}
	}
	return n, err
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	fileInfo, err := fromFile.Stat()
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	count := fileInfo.Size() - offset
	if limit > 0 && limit < count {
		count = limit
	}
	bar := pb.Start64(count)
	bar.Start()

	progressReader := &ProgressReader{
		Reader: fromFile,
		bar:    bar,
	}

	if limit > 0 {
		_, err = io.CopyN(toFile, progressReader, limit)
		if err != nil && errors.Is(err, io.EOF) {
			bar.Finish()
			return err
		}
	} else {
		_, err = io.Copy(toFile, progressReader)
		if err != nil {
			bar.Finish()
			return err
		}
	}
	bar.Finish()

	return nil
}
