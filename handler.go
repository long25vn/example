package main

import (
	"github.com/rs/xid"
	"os/exec"
)



func (cs *SliceContainer)HandlerRequest(content string, key string) error {

	localFilePath, fileName, err := CreateFile(content)
	if err != nil {
		return err
	}

	containerName := cs.GetFirstContainer()
	if containerName == "" {
		containerName = xid.New().String()
		_, err := CreateContainer(containerName)
		if err != nil {
			return err
		}
	} else {
		go cs.AppendNewContainerToSlice(key)
	}

	//copyFileCompile(localFilePath, containerName, containerRunDir string)

	err = copyFileCompile(localFilePath, containerName, "/home/dev")
	if err != nil {
		return err
	}

	//"docker", "exec", "-it", containerName, "go", "run", fileName})

	_, err = exec.Command("docker", "exec", "-it", containerName, "go", "run", fileName).Output()
	if err != nil {
		return err
	}
	return nil

}