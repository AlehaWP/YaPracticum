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

func (w *writer) Close() {
	w.file.Close()
}

func (r *reader) Close() {
	r.file.Close()
}

func newWriter(fileName string) (*writer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	file.Sync()
	if err != nil {
		return nil, errors.New("не удалось найти файл " + fileName)
	}
	defer file.Close()
	return &writer{
		file:    file,
		encoder: gob.NewEncoder(file),
	}, nil

}

func newReader(fileName string) (*reader, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.New("не удалось найти файл " + fileName)
	}
	defer file.Close()
	return &reader{
		file:    file,
		decoder: gob.NewDecoder(file),
	}, nil

}

var w *writer
var r *reader

func SaveURLSToFile(rep repository.URLRepo) {
	var err error
	w, err = newWriter(projectenv.Envs.OptionsFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.encoder.Encode(rep)

	// var buffer bytes.Buffer
	// gobEncoder := gob.NewEncoder(&buffer)

	// gobEncoder.Encode(rep)
	// fmt.Println(w.file.Name())
	// n, _ := w.file.WriteString("тест")
	// fmt.Println(n)

	w.Close()

	// gobDecoder := gob.NewDecoder(&buffer)
	// gobEncoder.Encode(rep)
	// gobDecoder.Decode(&rep)
	// fmt.Println("Decode rep", rep)
}

func ReadURLSFromFile(rep *repository.URLRepo) {
	r.decoder.Decode(rep)
	r.Close()
}

func init() {
	var err error

	r, err = newReader(projectenv.Envs.OptionsFileName)
	if err != nil {
		fmt.Println(err.Error())
	}

}
