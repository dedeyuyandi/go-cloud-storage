package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

type storageUploadFile struct {
	StorageClient *storage.Client
}

func NewStorage() (StorageUploadFile, error) {
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	return &storageUploadFile{
		StorageClient: storageClient,
	}, nil
}

func (os *storageUploadFile) Close() error {
	if os.StorageClient != nil {
		if err := os.StorageClient.Close(); err != nil {
			return err
		}
		os.StorageClient = nil
	}
	return nil
}
