package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var db = []string{"db1", "db2", "db3", "db4", "db5"}
var result = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(db); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(t0))

}

func dbCall(i int) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(i)
	log()
	wg.Done()
}

func save(i int) {
	m.Lock()
	result = append(result, db[i])
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("curent value of result: %v\n", result)
	m.RUnlock()
}
