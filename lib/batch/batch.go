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
	ch := make(chan struct{}, pool)
	//res = make([]user, 0, n)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(j int64) {
			temp := getOne(j)
			mu.Lock()
			<-ch
			res = append(res, temp)
			mu.Unlock()
			wg.Done()
		}(int64(i))
	}
	close(ch)
	wg.Wait()
	return res
}
