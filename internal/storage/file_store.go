package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/fromjyce/Nublink/internal/crypto"
	"github.com/fromjyce/Nublink/pkg/models"
	"github.com/google/uuid"
)

type FileStorage struct {
	basePath string
}

func NewFileStorage(basePath string) *FileStorage {
	return &FileStorage{basePath: basePath}
}

func (fs *FileStorage) StoreFile(filePath string, key []byte) string {
	id := uuid.New().String()
	storagePath := filepath.Join(fs.basePath, id)
	metaPath := filepath.Join(fs.basePath, id+".meta")

	// Encrypt and store the file
	if err := crypto.EncryptFile(key, filePath, storagePath); err != nil {
		panic(err)
	}

	// Store metadata
	meta := models.FileMetadata{
		ID:        id,
		Path:      storagePath,
		Key:       key,
		CreatedAt: time.Now(),
	}

	metaData, err := json.Marshal(meta)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(metaPath, metaData, 0600); err != nil {
		panic(err)
	}

	return id
}

func (fs *FileStorage) GetFile(id string) (*models.FileMetadata, []byte, error) {
	metaPath := filepath.Join(fs.basePath, id+".meta")
	storagePath := filepath.Join(fs.basePath, id)

	metaData, err := os.ReadFile(metaPath)
	if err != nil {
		return nil, nil, err
	}

	var meta models.FileMetadata
	if err := json.Unmarshal(metaData, &meta); err != nil {
		return nil, nil, err
	}

	fileData, err := crypto.DecryptFile(meta.Key, storagePath)
	if err != nil {
		return nil, nil, err
	}

	return &meta, fileData, nil
}

func (fs *FileStorage) DeleteFile(id string) error {
	metaPath := filepath.Join(fs.basePath, id+".meta")
	storagePath := filepath.Join(fs.basePath, id)

	if err := os.Remove(metaPath); err != nil {
		return err
	}
	return os.Remove(storagePath)
}
