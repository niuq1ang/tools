package redislock

import (
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func getRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		MaxActive:   1024,
		IdleTimeout: time.Duration(180) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("127.0.0.1:6379"))
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", 1)
			return c, nil
		},
	}
}

/*
	单请求，加锁解锁情况
*/
func TestRedisLock(t *testing.T) {
	pool := getRedisPool()
	key := "Project"
	doDemoRequst(pool, key, "JhWrCvGN")
}

/*
	多请求，竞争锁情况
*/
func TestMutilRedisLock(t *testing.T) {
	pool := getRedisPool()
	key := "Project"
	go doDemoRequst(pool, key, "JhWrCvGN")
	go doDemoRequst(pool, key, "DAyGEN19")
	go doDemoRequst(pool, key, "X7QCgzVj")

	time.Sleep(time.Second * 50)
}

func doDemoRequst(pool *redis.Pool, key, requestID string) {
	c := pool.Get()
	defer c.Close()
	lock, err := NewRedisSimpleLock(c, key, requestID, 15000, true)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer func() {
		lock.UnLock()
		fmt.Println(fmt.Sprintf("key:%s:requestID:%s-unlock!", key, requestID))
		ttl, err := getKeyExpire(c, key)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		if ttl == -2 {
			fmt.Println("success!!!")
		} else {
			fmt.Println("failure!!!, ttl:", ttl)
		}
	}()
	fmt.Println(fmt.Sprintf("key:%s:requestID:%s-lock!", key, requestID))
	ttl, err := getKeyExpire(c, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fmt.Sprintf("key:%s:requestID:%s-ttl:%d", key, requestID, ttl))
	// sleep
	fmt.Println("sleep 5 sesond unlock")
	time.Sleep(time.Second * 5)
}

func getKeyExpire(c redis.Conn, key string) (pexpire int, err error) {
	queryKey := fmt.Sprintf("simple_lock-%s", key) //方法里加了前缀
	pexpire, err = redis.Int(c.Do("PTTL", queryKey))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	return
}

/*
	增加锁过期时间
*/
func TestUpdateExpireTime(t *testing.T) {
	pool := getRedisPool()
	c := pool.Get()
	defer c.Close()
	key := "update_drafting"
	requestID := "2sks9a6X"
	time := int64(15000)
	lock, err := NewRedisSimpleLock(c, key, requestID, time, true)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("lock!!!")
	defer func() {
		lock.UnLock()
		fmt.Println("unlock!!!")
	}()
	pexpire, requestID, err := GetRedisSimpleInfo(c, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fmt.Sprintf("key:%s, requestID:%s, expire time:%d", key, requestID, pexpire))
	err = ExpireRedisSimpleLock(c, key, requestID, 20000)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	pexpire, requestID, err = GetRedisSimpleInfo(c, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fmt.Sprintf("key:%s, requestID:%s, expire time:%d", key, requestID, pexpire))
	if pexpire > time {
		fmt.Println("success!!!")
	} else {
		fmt.Println("failure!!!, ttl:", pexpire)
	}
}
