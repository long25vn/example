package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	gocache "github.com/patrickmn/go-cache"
)

type SliceContainer struct {
	sync.RWMutex
	Cache *gocache.Cache
	items []string
}

func (cs *SliceContainer) Append(item string) string {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
	first := cs.items[0]
	cs.items = cs.items[1:]
	fmt.Println(first)
	return first
}

func (cs *SliceContainer) GetFirstContainer(key string) string {
	cs.Lock()
	defer cs.Unlock()
	fmt.Println("len SliceContainer", len(cs.items))

	var slice  *SliceContainer
	value, found := cs.Cache.Get(key)
	if found {
		slice = value.(*SliceContainer)
		fmt.Println("slice ", slice)
	}

	if len(slice.items) > 0 {
		first := slice.items[0]
		slice.items = slice.items[1:]
		cs.Cache.Set(key, slice, gocache.DefaultExpiration)
		return first
	}
	return ""
}

func main() {
	conf := new(SliceContainer)

	conf.Cache = gocache.New(5*time.Minute, 10*time.Minute)
	key := "arrayContainer"

	conf.InitContainer(key)

	content := `
        package main
        import "fmt"
        func main() {
            sum := 0
            for i := 0; i < 10; i++ {
                sum += i
            }
            fmt.Println(sum)
        }`

	fmt.Println("-------")
	fmt.Println(conf.items)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		conf.HandlerRequest(content, key)
	})
	r.GET("/pong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run()
}