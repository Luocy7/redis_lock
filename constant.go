package redis_lock

import "time"

const (
	// KeyPrefix is the prefix of the lock key
	KeyPrefix = "redis_lock:"
	// DefaultLockExpire is the default expire time of the lock
	DefaultLockExpire = 8 * time.Second
	// DefaultBlockWaitingTime is the default block waiting time
	DefaultBlockWaitingTime = 5 * time.Second
	// WatchDogWorkStep is the time interval of the watchdog
	WatchDogWorkStep = 10 * time.Second
)
