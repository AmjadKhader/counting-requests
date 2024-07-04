package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type Storage interface {
	Save(data []time.Time)
	Load() []time.Time
}

type FileStorage struct {
	filename string
	mutex    sync.Mutex
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (fileStorage *FileStorage) Save(data []time.Time) {
	fileStorage.mutex.Lock()
	defer fileStorage.mutex.Unlock()
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("Could not marshal data: %s\n", err)
	}
	err = ioutil.WriteFile(fileStorage.filename, file, 0644)
	if err != nil {
		log.Fatalf("Could not write file: %s\n", err)
	}
}

func (fileStorage *FileStorage) Load() []time.Time {
	fileStorage.mutex.Lock()
	defer fileStorage.mutex.Unlock()
	var data []time.Time
	file, err := os.Open(fileStorage.filename)
	if err != nil {
		if os.IsNotExist(err) {
			fileStorage.Save([]time.Time{})
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
