package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (cs *SliceContainer) GetFirstContainer() string {
	cs.Lock()
	defer cs.Unlock()

	if len(cs.items) > 0 {
		first := cs.items[0]
		cs.items = cs.items[1:]
		return first
	}
	return ""
}

func main() {
	conf := new(SliceContainer)

	conf.Cache = gocache.New(5*time.Minute, 10*time.Minute)
	key := "arrayContainer"

	InitContainer(key, conf.Cache)

	content := `
        package main\n
        import "fmt"\n
        func main() {\n
            sum := 0\n
            for i := 0; i < 10; i++ {\n
                sum += i\n
            }\n
            fmt.Println(sum)\n
        }`

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		conf.HandlerRequest(content, key)
	})
	r.Run()
}