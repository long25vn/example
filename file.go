package main

import (
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"os/exec"
)

// Trả về file name đã được khởi tạo
func CreateFile(content string) (string, string, error) {
	var localFolderPath, localFilePath, fileName string

	folderID := xid.New().String()
	localFolderPath = "." + "/" + "go" + "/" + folderID


	localFilePath = localFolderPath + "/" + fileName

	if _, err := os.Stat(localFolderPath); os.IsNotExist(err) {
		os.MkdirAll(localFolderPath, os.FileMode(0777))
	}

	file, err := os.Create(localFilePath)
	if err != nil {
		return localFilePath, fileName, err
	}
	defer file.Close()

	code := []byte(content)
	err = ioutil.WriteFile(localFilePath, code, 0644)

	return localFilePath, fileName, err
}

// Copy file vào trong container
func copyFileCompile(localFilePath, containerName, containerRunDir string) error {
	_, err := exec.Command("docker", "cp", localFilePath, fmt.Sprintf("%s:%s", containerName, containerRunDir)).CombinedOutput()
	return err
}
