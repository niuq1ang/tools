package redislock

import (
	"fmt"
	"time"

	"github.com/bangwork/bang-api/app/utils/errors"
	"github.com/bangwork/bang-api/app/utils/log"
	"github.com/garyburd/redigo/redis"
)

var (
	redisLockByKeyRetryRate = 5
	minDelay                = redisLockByKeyRetryRate
)

const (
	SimpleLockNotFoundExpire = -2
)

type SimpleLock struct {
	c     redis.Conn
	k     string
	id    string
	px    int64
	block bool
}

/*
	通过key和requestID加锁
	block=true时，如果锁已被占用则阻塞至解锁
	block=false时，如果锁已被占用则返回异常
*/
func NewRedisSimpleLock(c redis.Conn, key, requestID string, pexpire int64, block bool) (*SimpleLock, error) {
	sl := new(SimpleLock)
	sl.c = c
	sl.k = key
	sl.id = requestID
	sl.px = pexpire
	sl.block = block
	err := sl.getLockToken()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return sl, nil
}

func lockKey(s string) string {
	return fmt.Sprintf("simple_lock-%s", s)
}

func (sl *SimpleLock) getLockToken() error {
	c := sl.c
	key := lockKey(sl.k)
	lockF := func() (interface{}, error) {
		return c.Do("SET", key, sl.id, "PX", sl.px, "NX")
	}
	reply, err := lockF()
	if err != nil {
		return errors.Redis(err)
	}
	if reply == "OK" {
		return nil
	}
	if !sl.block {
		return fmt.Errorf("key:%s, requestID:%s lock bussy", key, sl.id)
	}
	ttl, err := redis.Int(c.Do("PTTL", key))
	if err != nil && err != redis.ErrNil {
		return errors.Redis(err)
	}
	if ttl > 0 {
		delay := ttl / redisLockByKeyRetryRate
		if delay == 0 {
			delay = minDelay
		}
		for i := 0; i <= redisLockByKeyRetryRate; i++ {
			time.Sleep(time.Millisecond * time.Duration(delay))
			reply, err := lockF()
			if err != nil {
				return errors.Redis(err)
			}
			if reply == "OK" {
				return nil
			}
		}
	}
	return nil
}

func (sl *SimpleLock) UnLock() {
	err := UnLockRedisSimpleLock(sl.c, sl.k, sl.id)
	if err != nil {
		log.Warn("UnLockRedisSimpleLock error:%v", err)
	}
}

/*
	通过key和requestID解锁
*/
func UnLockRedisSimpleLock(c redis.Conn, key, requestID string) error {
	key = lockKey(key)
	getScript := redis.NewScript(1, "if redis.call('GET', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end")
	_, err := getScript.Do(c, key, requestID)
	if err != nil {
		return errors.Redis(err)
	}
	return nil
}

/*
	通过key和requestID延长加锁时间
*/
func ExpireRedisSimpleLock(c redis.Conn, key, requestID string, pexpire int64) error {
	key = lockKey(key)
	getScript := redis.NewScript(1, "if redis.call('GET', KEYS[1]) == ARGV[1] then return redis.call('PEXPIRE', KEYS[1], ARGV[2]) else return 0 end")
	_, err := getScript.Do(c, key, requestID, pexpire)
	if err != nil {
		return errors.Redis(err)
	}
	return nil
}

/*
	通过key获取锁信息
*/
func GetRedisSimpleInfo(c redis.Conn, key string) (int64, string, error) {
	key = lockKey(key)
	pexpire, err := redis.Int64(c.Do("PTTL", key))
	if err == redis.ErrNil {
		return pexpire, "", nil
	} else if err != nil {
		return pexpire, "", errors.Redis(err)
	}
	value, err := redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		return pexpire, "", nil
	} else if err != nil {
		return pexpire, "", errors.Redis(err)
	}
	return pexpire, value, nil
}
