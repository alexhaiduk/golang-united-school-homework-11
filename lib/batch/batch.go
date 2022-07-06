package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan []user, pool)
	//res = make([]user, 0)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		go func(j int64) {
			//res = append(res, getOne(j))
			mu.Lock()
			ch <- append(<-ch, getOne(j))
			mu.Unlock()
		}(int64(i))
		wg.Done()
	}
	wg.Wait()
	return <-ch
}
