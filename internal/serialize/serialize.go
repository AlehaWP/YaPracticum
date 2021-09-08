package serialize

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"

	"github.com/AlehaWP/YaPracticum.git/internal/projectenv"
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
	defer file.Close()
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
	defer file.Close()
	return &reader{
		file:    file,
		decoder: gob.NewDecoder(file),
	}, nil

}

var w *writer
var r *reader

func SaveURLSToFile(rep *repository.URLRepo) {
	w.encoder.Encode(rep)
}

func ReadURLSFromFile(rep *repository.URLRepo) {
	r.decoder.Decode(rep)
}

func init() {
	var err error
	w, err = newWriter(projectenv.Envs.OptionsFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	r, err = newReader(projectenv.Envs.OptionsFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}
