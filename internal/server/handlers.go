package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (s *Server) downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/download/"):]
	if id == "" {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	meta, data, err := s.fileStore.GetFile(id)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Check expiration
	if meta.ExpiresAt != nil && time.Now().After(*meta.ExpiresAt) {
		http.Error(w, "Link has expired", http.StatusGone)
		s.fileStore.DeleteFile(id)
		return
	}

	// Check one-time download
	if r.URL.Query().Get("once") == "true" || meta.OneTime {
		if meta.Accessed {
			http.Error(w, "Link has already been used", http.StatusGone)
			s.fileStore.DeleteFile(id)
			return
		}
		meta.Accessed = true
	}

	// Update metadata if needed
	if meta.Accessed {
		metaData, err := json.Marshal(meta)
		if err == nil {
			metaPath := fmt.Sprintf("%s/%s.meta", s.cfg.StoragePath, id)
			_ = os.WriteFile(metaPath, metaData, 0600)
		}
	}

	// Serve the file
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filepath.Base(meta.Path)))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)

	// Delete if one-time
	if r.URL.Query().Get("once") == "true" || meta.OneTime {
		go func() {
			time.Sleep(1 * time.Second) // Give time for download to complete
			s.fileStore.DeleteFile(id)
		}()
	}
}
