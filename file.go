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


func CreateFileInContainer(containerName, content string) (string, error) {
	//var localFolderPath, localFilePath string
	fmt.Println("CreateFileInContainer")
	xid := xid.New().String()
	fileName := xid + ".go"
	//y := fmt.Sprint(content)
	go func(){
		x1 := fmt.Sprintln("mkdir", "-p", "duong/duong2/duong3")
		out2, err := exec.Command("docker", "exec", containerName, "sh", "-c", x1).Output()
		if err != nil {
			fmt.Println("err CreateFileInContainer", err)
		}
		fmt.Println("out", string(out2))
	}()
	go func(){
		x1 := fmt.Sprintln("mkdir", "-p", "duong/duong2/duong4")
		out2, err := exec.Command("docker", "exec", containerName, "sh", "-c", x1).Output()
		if err != nil {
			fmt.Println("err CreateFileInContainer", err)
		}
		fmt.Println("out", string(out2))
	}()
	time.Sleep(time.Millisecond*200)
	x := fmt.Sprintln("echo", "'", content, "'", ">", "duong/duong2/duong3/main.go")
	out, err := exec.Command("docker", "exec", containerName, "sh", "-c", x).Output()
	if err != nil {
		fmt.Println("err CreateFileInContainer", err)
	}
	fmt.Println("out", string(out))


	return fileName, nil
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
