package main

import (
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"github.com/rs/xid"
	"os/exec"
	"regexp"
)
func InitContainer(key string, c *gocache.Cache) (error) {
	xid := xid.New().String()

	var slice *SliceContainer = &SliceContainer{}
	value, found := c.Get(key)
	if found {
		slice = value.(*SliceContainer)
	}


	numberContainer := len(slice.items)
	for numberContainer < 5 {
		isExist ,err := CreateContainer(xid)
		if !isExist && err == nil {
			slice.items = append(slice.items, xid)
			c.Set(key, slice, gocache.DefaultExpiration)
			numberContainer += 1
			fmt.Println("Created Container")
			fmt.Println(slice)
		}
	}

	return nil
}


//func ManageContainer(key string, c *gocache.Cache) (error) {
//	xid := xid.New().String()
//
//
//	var slice *SliceContainer
//	value, found := c.Get(key)
//	if found {
//		slice = value.(*SliceContainer)
//	}
//
//
//
//	numberContainer := len(slice.items)
//	for len(slice.items) < 5 {
//		err := CreateContainer(xid)
//		if err != nil {
//			numberContainer += 1
//			slice.items = append(slice.items, xid)
//			c.Set(key, slice, gocache.DefaultExpiration)
//		}
//	}
//
//	return nil
//}

func  (cs *SliceContainer)AppendNewContainerToSlice(key string) (error) {
	var slice *SliceContainer
	value, found := cs.Cache.Get(key)
	if found {
		slice = value.(*SliceContainer)
	}

	xid := xid.New().String()
	isExist, err := CreateContainer(xid)
	if err != nil {
		return nil
	}
	if !isExist {
		slice.items = append(slice.items, xid)
		cs.Cache.Set(key, slice, gocache.DefaultExpiration)
	}
	return nil
}


// Nếu container đã tồn tại trước đó, trả về true
func CreateContainer(containerName string) (bool,error) {
	isExist := isContainerExist(containerName)
	if isExist {
		return true, nil
	}

	_, err := exec.Command("docker", "run", "-id", "--rm", "--name", containerName, "compiler-go").Output()
	if err != nil {
		return false, err
	}
	return false, nil
}

func isContainerExist(containerName string) (isExist bool) {
	out, _ := exec.Command("docker", "inspect", "--format=\"{{.Name}}\"", containerName).CombinedOutput()

	regexContainerExist, _ := regexp.Compile("No such object: " + containerName)
	isExist = !regexContainerExist.MatchString(string(out))

	return isExist
}

func RemoveContainer(containerName string) (string, error) {
	_, err := exec.Command("docker", "kill", containerName).CombinedOutput()
	return containerName, err
}