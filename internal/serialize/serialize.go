package serialize

import (
	"encoding/gob"
	"errors"
	"os"

	"github.com/AlehaWP/YaPracticum.git/internal/repository"
)

type reader struct {
	file    *os.File
	decoder *gob.Decoder
}

type writer struct {
	file    *os.File
	encoder *gob.Encoder
}

func newWriter(fileName string) (*writer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return nil, errors.New("не удалось найти файл")
	}
	return &writer{
		file:    file,
		encoder: gob.NewEncoder(file),
	}, nil

}

func newReader(fileName string) (*reader, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, errors.New("не удалось найти файл")
	}
	return &reader{
		file:    file,
		decoder: gob.NewDecoder(file),
	}, nil

}

var w *writer
var r *reader

func Init(fileName string) error {
	var err error
	w, err = newWriter(fileName)
	if err != nil {
		return err
	}
	r, err = newReader(fileName)
	if err != nil {
		return err
	}

	return nil
}

func SaveURLSToFile(rep *repository.URLRepo) {
	w.encoder.Encode(rep)
}

func ReadURLSFromFile(rep *repository.URLRepo) {
	r.decoder.Decode(rep)
}
