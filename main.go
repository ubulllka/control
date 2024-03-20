package main

import (
	"context"
	"sync"
	"task/config"
	"task/fcontrol"
	"task/redis"
	"time"
)

func main() {

	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	url := config.GetConf().Redis.URL
	client, err := redis.InitDB(url)
	if err != nil {
		panic(err)
	}

	fc := fcontrol.Fcontrol{
		Max:     100,
		DurTime: 3 * time.Second,
		Redis:   client,
	}

	var wg sync.WaitGroup

	ctx := context.Background()

	wg.Add(2)

	go func(fc *fcontrol.Fcontrol) {
		var q bool
		var i int64 = 1
		for {
			if q {
				defer wg.Done()
				return
			} else {
				flag, err := fc.Check(ctx, i)
				if err != nil {
					panic(err)
				}
				q = flag
			}
			i++
		}
	}(&fc)

	go func(fc *fcontrol.Fcontrol) {
		var q bool
		var i int64 = 1
		for {
			if q {
				defer wg.Done()
				return
			} else {
				flag, err := fc.Check(ctx, i)
				if err != nil {
					panic(err)
				}
				q = flag
			}
			i++
		}
	}(&fc)

	wg.Wait()
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
