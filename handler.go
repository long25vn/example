package main

import (
	"fmt"
	"github.com/rs/xid"
	"os/exec"
	"time"
)



func (cs *SliceContainer)HandlerRequest(content string, key string) error {

	var containerName string
	//ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// Tạo file trên local từ code
	//go func() {
	//	var err error
	//	localFilePath, fileName, err = CreateFile(content)
	//	if err != nil {
	//		ch1 <- err.Error()
	//	}
	//	ch1 <- "done"
	//}()


	go func() {
		containerName = cs.GetFirstContainer(key)
		fmt.Println("containerName hadler ", containerName)
		if containerName == "" {
			startTime := time.Now()
			containerName = xid.New().String()
			_, err := CreateContainer(containerName)
			fmt.Println("time tao container ", time.Since(startTime))
			if err != nil {
				panic(err)
			}
			ch2 <- "done"
		} else {
			ch2 <- "done"
			go cs.AppendNewContainerToSlice(key)
		}
	}()

	//result1 := <-ch1
	result2 := <-ch2
	if result2 == "done" {
		//err := copyFileCompile(localFilePath, containerName, "/home/dev")
		//if err != nil {
		//	return err
		//}
		fileName, err := CreateFileInContainer(containerName, content)
		if err != nil {
			fmt.Println("err", err)
		}

		out, err := exec.Command("docker", "exec", containerName, "go", "run", fileName).Output()
		if err != nil {
			return err
		}
		fmt.Println("output", string(out))
	}

	return nil
}