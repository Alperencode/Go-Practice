package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("\nTime passed: ", time.Since(t0))
}

func dbCall(index int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nData from db is: %v", dbData[index])
	wg.Done()
}
