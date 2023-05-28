package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	folderPath := "enc"

	var oldExtension, newExtension string

	fmt.Print("Masukkan ekstensi yang ingin diubah: ")
	fmt.Scanln(&oldExtension)
	fmt.Print("Masukkan ekstensi baru: ")
	fmt.Scanln(&newExtension)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == oldExtension {
				newPath := path[:len(path)-len(oldExtension)] + newExtension
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}
				fmt.Printf("File %s diubah menjadi %s\n", path, newPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Terjadi kesalahan: %v\n", err)
	}
}
