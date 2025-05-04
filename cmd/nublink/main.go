package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fromjyce/Nublink/internal/crypto"
	"github.com/fromjyce/Nublink/internal/server"
	"github.com/fromjyce/Nublink/internal/storage"
	"github.com/fromjyce/Nublink/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nublink",
	Short: "Secure file sharing with self-destruct links",
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a file securely",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		expire, _ := cmd.Flags().GetString("expire")
		once, _ := cmd.Flags().GetBool("once")

		cfg := config.LoadConfig()
		key := crypto.GenerateKey()
		fileStore := storage.NewFileStorage(cfg.StoragePath)
		fileID := fileStore.StoreFile(filePath, key)

		srv := server.NewServer(cfg, fileStore)
		go srv.Start()

		url := fmt.Sprintf("https://localhost:%d/download/%s", cfg.Port, fileID)
		if once {
			url += "?once=true"
		} else if expire != "" {
			duration, _ := time.ParseDuration(expire)
			url += fmt.Sprintf("?expire=%d", time.Now().Add(duration).Unix())
		}

		fmt.Printf("Share this link: %s\n", url)
		fmt.Println("Press Ctrl+C to stop the server")
		select {}
	},
}

func init() {
	shareCmd.Flags().StringP("file", "f", "", "File to share (required)")
	shareCmd.MarkFlagRequired("file")
	shareCmd.Flags().StringP("expire", "e", "", "Expiration duration (e.g., 1h, 30m)")
	shareCmd.Flags().BoolP("once", "o", false, "One-time download")
	rootCmd.AddCommand(shareCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
