package mutexmangager

import (
	"fmt"
	"seeder/logger"
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex

func TestMutexManager(t *testing.T) {
	fmt.Println(SeederLogger.Author)
	go func() {
		mu.Lock()
		fmt.Printf("go 1 ")
		mu.Unlock()

		go func() {
			fmt.Printf("enter 2 ")

			mu.Lock()
			time.Sleep(time.Second * 1)
			fmt.Printf("go 2 ")
			mu.Unlock()
		}()
		go func() {
			fmt.Printf("enter 3 ")

			mu.Lock()
			time.Sleep(time.Second * 1)
			fmt.Printf("go 3 ")
			mu.Unlock()
		}()
		go func() {
			fmt.Printf("enter 4 ")

			mu.Lock()
			time.Sleep(time.Second * 1)
			fmt.Printf("go 4 ")
			mu.Unlock()
		}()
	}()

	time.Sleep(time.Second * 4)
	// mm := NewMutexManager()

}
