package models

import (
	"time"
)

type FileMetadata struct {
	ID        string
	Path      string
	Key       []byte
	CreatedAt time.Time
	ExpiresAt *time.Time
	OneTime   bool
	Accessed  bool
}
