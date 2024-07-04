package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Storage interface {
	Save(data []time.Time)
	Load() []time.Time
}

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (f *FileStorage) Save(data []time.Time) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("Could not marshal data: %s\n", err)
	}
	err = ioutil.WriteFile(f.filename, file, 0644)
	if err != nil {
		log.Fatalf("Could not write file: %s\n", err)
	}
}

func (f *FileStorage) Load() []time.Time {
	var data []time.Time
	file, err := os.Open(f.filename)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, create it with an initial empty array
			f.Save([]time.Time{})
			return data
		} else {
			log.Fatalf("Could not open file: %s\n", err)
		}
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Could not read file: %s\n", err)
	}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Could not unmarshal data: %s\n", err)
	}
	return data
}
