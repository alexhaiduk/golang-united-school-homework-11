// goroutine1 project goroutine1.go
package main

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	ID int64
}

func main() {
	/*	var wg sync.WaitGroup
		var mux sync.Mutex
		var t []user
		//var o = user{ID: 1}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			//ch <- "hello, world!"
			//ch <- i
			go func(j int) {
				mux.Lock()
				defer mux.Unlock()
				fmt.Println(t)
				t = append(t, getOne(int64(j)))
				//o.ID++
				//time.Sleep(1 * time.Second)
				wg.Done()
			}(i)
		}

		wg.Wait()*/
	fmt.Println(getBatch(30, 5))
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan user, pool)
	res = make([]user, 0, n)

	for i := 0; i < int(n); i++ {
		wg.Add(1)
		//user1 := getOne(int64(i))
		ch <- getOne(int64(i))
		go func() {
			mu.Lock()
			defer mu.Unlock()
			//temp = append(temp, getOne(j))
			res = append(res, <-ch)
			wg.Done()
		}()
	}
	wg.Wait()
	//res = temp
	return res
}
