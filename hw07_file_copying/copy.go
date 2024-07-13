package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Открываем файл
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	//закрываем после копирования
	defer fromFile.Close()

	// получаем инфо о файле
	fileInfo, err := fromFile.Stat()
	if err != nil {
		return err
	}

	// fmt.Println(fileInfo.Size())
	// os.Exit(0)

	// проверка что это обычный файл
	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	// если указатель чтения выходит за границы файла - ошибка
	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	// создаем файл куда копируем данные
	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	// Перемещаем указатель чтения из исходного файла на offset
	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	// count := int(fileInfo.Size()
	// bar := pb.StartNew(count)
	// for i := 0; i < count; i++ {
	// 	bar.Increment()
	// 	time.Sleep(time.Millisecond)
	// }
	// bar.Finish()

	// Копируем данные
	if limit > 0 {
		_, err = io.CopyN(toFile, fromFile, limit)
		if err != nil && err != io.EOF {
			return err
		}
	} else {
		_, err = io.Copy(toFile, fromFile)
		if err != nil {
			return err
		}
	}

	return nil
}
