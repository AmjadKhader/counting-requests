package storage

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper function to normalize time slices for comparison
func normalizeTimes(timestamps []time.Time) []time.Time {
	var normalized []time.Time
	for _, ts := range timestamps {
		normalized = append(normalized, ts.UTC())
	}
	return normalized
}

func TestFileStorageSaveLoad(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "filestorage_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	store := NewFileStorage(tmpFile.Name())

	timestamps := []time.Time{time.Now().UTC()}
	store.Save(timestamps)

	loadedTimestamps := store.Load()
	assert.Equal(t, normalizeTimes(timestamps), normalizeTimes(loadedTimestamps))
}

func TestFileStorageLoadNonExistentFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "filestorage_test")
	assert.NoError(t, err)
	os.Remove(tmpFile.Name()) // Remove the file to simulate non-existence

	store := NewFileStorage(tmpFile.Name())

	loadedTimestamps := store.Load()
	assert.Equal(t, []time.Time(nil), loadedTimestamps)

	// Check if the file is created
	_, err = os.Stat(tmpFile.Name())
	assert.False(t, os.IsNotExist(err))
}

func TestFileStorageSaveCreatesFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "filestorage_test")
	assert.NoError(t, err)
	os.Remove(tmpFile.Name()) // Remove the file to simulate non-existence

	store := NewFileStorage(tmpFile.Name())

	timestamps := []time.Time{time.Now().UTC()}
	store.Save(timestamps)

	// Check if the file is created
	_, err = os.Stat(tmpFile.Name())
	assert.False(t, os.IsNotExist(err))

	loadedTimestamps := store.Load()
	assert.Equal(t, normalizeTimes(timestamps), normalizeTimes(loadedTimestamps))
}
