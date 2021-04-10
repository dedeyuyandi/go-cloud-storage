package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/dedeyuyandi/go-cloud-storage/constanta"
)

func (s *storageUploadFile) UploadFile(ctx context.Context, req interface{}) (string, error) {
	f, err := os.Open("notes.txt")
	if err != nil {
		return "", fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := s.StorageClient.Bucket(constanta.BucketName).Object(constanta.ObjectName).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	w := io.Writer(f)
	fmt.Fprintf(w, "Blob %v uploaded.\n", constanta.ObjectName)

	return "", nil
}
