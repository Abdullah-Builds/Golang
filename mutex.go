package main

import (
	"fmt"
	"sync"
)

type counter struct{
	count int 
	mutex sync.Mutex
}

func(c *counter) update(wg* sync.WaitGroup){
	defer func(){
		wg.Done() 
		c.mutex.Unlock()}()

	c.mutex.Lock()
	c.count ++
}

func main(){
    var Obj = counter{count : 0}
	var wg sync.WaitGroup
    
	for i:=0; i<100; i++ {
		wg.Add(1)
        go Obj.update(&wg)
		fmt.Println(Obj.count)
	}

	wg.Wait()
}