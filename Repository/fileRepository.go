package Repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rickdana/fizzbuzzApi/Model"
	"io/ioutil"
	"log"
	"os"
)

type FileRepository struct {
	config Config
}

func NewFileRepository(config Config) *FileRepository {
	if _, err := CreateFile(config.GetConfig()); err != nil {
		panic(err)
	}
	return &FileRepository{config: config}
}

func (fr *FileRepository) Save(entity interface{}) (err error) {
	f, err := os.OpenFile(fr.config.GetConfig(), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(entity)
	n, err := f.Write(b)

	if err != nil {
		return err
	}

	n, err = f.WriteString("\n")

	if err != nil {
		return err
	}
	log.Print(n)
	return nil
}

func (fr *FileRepository) FindAll() ([]Model.FizzBuzz, error) {
	file, err := ioutil.ReadFile(fr.config.GetConfig())

	if err != nil {
		return nil, err
	}

	s := string(file)

	fizzBuzzDtos, err := Model.FizzBuzzDtoFactory(s)

	if err != nil {
		return nil, err
	}

	return fizzBuzzDtos, nil
}

func CreateFile(filePath string) (string, error) {
	var _, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		var file, err = os.Create(filePath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Unable to create file %v", filePath))
		}
		defer file.Close()
	}

	return filePath, nil
}
