package redis

import (
	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	letters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lockCommand = `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
            return "OK"
        else
            return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
        end`
	unlockCommand = `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end`
	randomLen       = 16
	tolerance       = 500
	millisPerSecond = 1000
)

type RedisLock struct {
	store   *redigo.Conn
	seconds uint32
	key     string
	value   string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRedisLock(store *redigo.Conn, key, value string) *RedisLock {
	if value == "" {
		value = randomStr(randomLen)
	}
	return &RedisLock{
		store: store,
		key:   key,
		value: value,
	}
}

func (rl *RedisLock) Lock() (bool, error) {
	return rl.LockCtx(nil)
}

func (rl *RedisLock) LockCtx(ctx *gin.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	script := redigo.NewScript(1, lockCommand)
	resp, err := script.DoContext(ctx, *rl.store, rl.key, rl.value, strconv.Itoa(int(seconds)*millisPerSecond+tolerance))
	if err == redigo.ErrNil {
		return false, nil
	} else if err != nil {
		return false, err
	} else if resp == nil {
		return false, nil
	}
	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}
	return false, nil
}

func (rl *RedisLock) Unlock() (bool, error) {
	return rl.UnlockCtx(nil)
}

func (rl *RedisLock) UnlockCtx(ctx *gin.Context) (bool, error) {
	script := redigo.NewScript(1, unlockCommand)
	resp, err := script.DoContext(ctx, *rl.store, rl.key, rl.value)
	if err != nil {
		return false, err
	}
	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}
	return reply == 1, nil
}

func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}

func randomStr(n int) string {
	b := make([]byte, 0, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
