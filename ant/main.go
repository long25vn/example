
package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
)

type ConcurrentSlice struct {
	sync.RWMutex
	items []int
}

func (cs *ConcurrentSlice) Append(item int) int {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
	first := cs.items[0]
	cs.items = cs.items[1:]
	fmt.Println(first)
	return first
}
var sum int32

func myFunc(i interface{}) error {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	return nil
}

func main() {

	defer ants.Release()

	//var x ConcurrentSlice
	runTimes := 100
	var wg sync.WaitGroup

	// use the common poo
	//for i := 0; i < runTimes; i++ {
	//	wg.Add(1)
	//	ants.Submit(func() error {
	//		x.Append(i)
	//		wg.Done()
	//		return nil
	//	})
	//}

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) error {
		myFunc(i)
		wg.Done()
		return nil
	})
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Serve(int32(i))
	}
	wg.Wait()

	//fmt.Println(x)

}