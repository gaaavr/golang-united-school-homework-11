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
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	sem := make(chan struct{}, pool)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			defer wg.Done()
			defer func() {
				<-sem
			}()
			u := getOne(j)
			mu.Lock()
			res = append(res, u)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return res
}
