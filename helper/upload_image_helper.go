package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadImageHelper(fileHeader *multipart.FileHeader, uploadPath string) (string, error) {
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("error creating directory: %v", err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing destination file: %v\n", err)
			fmt.Printf("error closing file: %v\n", err)
		}
	}(file)

	uniqueFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	savePath := filepath.Join(uploadPath, uniqueFileName)

	dst, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			fmt.Printf("error closing destination file: %v\n", err)
		}
	}(dst)

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("error saving file: %v", err)
	}

	path := strings.ReplaceAll(savePath, "\\", "/")

	return path, nil
}
