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
	//ch := make(chan []user, pool)
	res = make([]user, 0, n)
	//wg.Add(int(pool))
	for i := 0; i < int(n); i++ {
		wg.Add(int(pool))
		//go func(j int64) {
		go func() {
			defer wg.Done()
			mu.Lock()
			//tmp := getOne(j)
			//res = append(res, tmp)
			res = append(res, getOne(int64(i)))
			//res = append(res, getOne(j))
			//ch <- append(res, getOne(int64(j)))
			//ch <- res
			mu.Unlock()
			//ch <- res
			//}(int64(i))
			//ch <- res
		}()
		//wg.Done()
	}
	wg.Wait()
	//return <-ch
	//res = <-ch
	//return res
	return res
}
