package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrWhileClosingFile      = errors.New("error while closing file, data was corrpted")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func() {
		if closeError := fileTo.Close(); closeError != nil {
			err = ErrWhileClosingFile
		}
	}()

	fs, err := fileFrom.Stat()
	if err != nil {
		// * программа может НЕ обрабатывать файлы, у которых неизвестна длина (например, /dev/urandom);
		return ErrUnsupportedFile
	}

	if offset > fs.Size() {
		// * offset больше, чем размер файла - невалидная ситуация;
		return ErrOffsetExceedsFileSize
	}

	fsOffset := fs.Size() - offset
	if limit == 0 || limit > fsOffset {
		limit = fsOffset
	}

	_, err = fileFrom.Seek(offset, 0)
	if err != nil {
		return err
	}

	bar := pb.Start64(limit)
	defer bar.Finish()
	var count int64
	for {
		time.Sleep(time.Microsecond)
		bar.Increment()
		wr, err := io.CopyN(fileTo, fileFrom, 1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("failed to copy files %w", err)
		}
		count += wr
		if count == limit {
			break
		}
	}

	return nil
}
