package main

import (
	"fmt"
	"github.com/rs/xid"
	"os/exec"
	"time"
)



func (cs *SliceContainer)HandlerRequest(content string, key string) error {
	//
	localFilePath, fileName, err := CreateFile(content)
	if err != nil {
		return err
	}
	//
	containerName := cs.GetFirstContainer()
	if containerName == "" {
		startTime := time.Now()
		containerName = xid.New().String()
		_, err := CreateContainer(containerName)
		fmt.Println("time.Since",time.Since(startTime))
		if err != nil {
			return err
		}
	} else {
		go cs.AppendNewContainerToSlice(key)
	}
	fmt.Println("containerName",containerName)

	err = copyFileCompile(localFilePath, containerName, "/home/dev")
	if err != nil {
		return err
	}

	out, err := exec.Command("docker", "exec", containerName, "go", "run", fileName).Output()
	if err != nil {
		return err
	}
	fmt.Println("output", string(out))
	return nil
}