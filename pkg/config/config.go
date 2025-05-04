package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Port        int
	StoragePath string
	TLSCert     string
	TLSKey      string
}

func LoadConfig() Config {
	home, _ := os.UserHomeDir()
	storagePath := filepath.Join(home, ".nublink", "files")
	_ = os.MkdirAll(storagePath, 0700)

	return Config{
		Port:        8443,
		StoragePath: storagePath,
		TLSCert:     filepath.Join(home, ".nublink", "cert.pem"),
		TLSKey:      filepath.Join(home, ".nublink", "key.pem"),
	}
}
