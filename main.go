package main

import (
	"fmt"

	"github.com/dedeyuyandi/go-cloud-storage/storage"
)

func main() {
	fmt.Println("main")
	_, err := storage.NewStorage()
	if err != nil {
		fmt.Println("Connection to storage service failed", err.Error())
	}
}
