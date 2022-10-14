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
	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`
	unlockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
	randomLen = 16
	// 默认超时时间，防止死锁
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

// A RedisLock is a redis lock.
type RedisLock struct {
	// redis客户端
	store *redigo.Conn
	// 超时时间
	seconds uint32
	// 锁key
	key string
	// 锁value，防止锁被别人获取到
	value string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRedisLock returns a RedisLock.
func NewRedisLock(store *redigo.Conn, key string, value string) *RedisLock {
	if len(value) == 0 {
		value = randomStr(randomLen)
	}
	return &RedisLock{
		store: store,
		key:   key,
		value: value,
	}
}

// Lock acquires the lock.
func (rl *RedisLock) Lock() (bool, error) {
	return rl.LockCtx(nil)
}

// LockCtx acquires the lock with the given ctx.
func (rl *RedisLock) LockCtx(ctx *gin.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	var script = redigo.NewScript(1, lockCommand)
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

// Unlock releases the lock.
func (rl *RedisLock) Unlock() (bool, error) {
	return rl.UnlockCtx(nil)
}

// UnlockCtx releases the lock with the given ctx.
func (rl *RedisLock) UnlockCtx(ctx *gin.Context) (bool, error) {
	var script = redigo.NewScript(1, unlockCommand)
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

// SetExpire sets the expiration.
// 需要注意的是需要在Acquire()之前调用
// 不然默认为500ms自动释放
func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
