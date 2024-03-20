package fcontrol

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
)

type Fcontrol struct {
	sync.RWMutex
	Max     int64
	DurTime time.Duration
	Redis   *redis.Client
}

func (f *Fcontrol) Check(ctx context.Context, userID int64) (bool, error) {
	f.Lock()
	defer f.Unlock()

	now := time.Now().Unix()
	nowStr := strconv.FormatInt(now, 10)
	startStr := strconv.FormatInt(now-int64(f.DurTime), 10)

	queryCount := f.Redis.ZCount(ctx, strconv.FormatInt(userID, 10), startStr, nowStr)
	count, err := queryCount.Result()
	if err != nil {
		return false, err
	}

	if count >= f.Max {
		return false, nil
	}

	queryAdd := f.Redis.ZAdd(ctx, strconv.FormatInt(userID, 10), redis.Z{Score: float64(now), Member: uuid.New().String()})
	_, err = queryAdd.Result()
	if err != nil {
		return false, err
	}

	expireCmd := f.Redis.Expire(ctx, strconv.FormatInt(userID, 10), f.DurTime)
	_, err = expireCmd.Result()
	if err != nil {
		return false, err
	}

	return true, nil
}
