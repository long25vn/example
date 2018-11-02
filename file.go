package main

import (
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// Trả về file name đã được khởi tạo
func CreateFile(content string) (string, string, error) {
	var localFolderPath, localFilePath, fileName string

	folderID := xid.New().String()
	localFolderPath = "." + "/" + "go" + "/" + folderID

	fileName = "main.go"
	localFilePath = localFolderPath + "/" + fileName

	timeCreateFolder := time.Now()
	if _, err := os.Stat(localFolderPath); os.IsNotExist(err) {
		os.MkdirAll(localFolderPath, os.FileMode(0777))
	}
	fmt.Println("timeCreateFolder", time.Since(timeCreateFolder))


	timeCreateFile := time.Now()
	file, err := os.Create(localFilePath)
	if err != nil {
		return localFilePath, fileName, err
	}
	defer file.Close()

	code := []byte(content)
	err = ioutil.WriteFile(localFilePath, code, 0644)
	fmt.Println("timeCreateFile", time.Since(timeCreateFile))

	return localFilePath, fileName, err
}

// Copy file vào trong container
func copyFileCompile(localFilePath, containerName, containerRunDir string) error {
	out, err := exec.Command("docker", "cp", localFilePath, fmt.Sprintf("%s:%s", containerName, containerRunDir)).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Copy File", string(out))
	return nil
}
