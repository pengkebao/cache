package cache

import (
	"errors"
	"time"
)

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}

var (
	Instance        Cache
	ErrCacheMiss    = errors.New("revel/cache: key not found")
	ErrNotStored    = errors.New("revel/cache: not stored")
	ErrInvalidValue = errors.New("revel/cache: invalid value")
)

func Get(key string) interface{} { return Instance.Get(key) }
func Set(key string, val interface{}, timeout time.Duration) error {
	return Instance.Set(key, val, timeout)
}
func IsExist(key string) bool { return Instance.IsExist(key) }
func Delete(key string) error { return Instance.Delete(key) }
