package cache

import "time"

type Memory struct {
}

func NewMemory() *Memory {
	return &Memory{}
}

func (c *Memory) Get(key string) interface{} {
	return nil
}
func (c *Memory) Set(key string, val interface{}, timeout time.Duration) error {
	return nil
}
func (c *Memory) IsExist(key string) bool {
	return false
}
func (c *Memory) Delete(key string) error {
	return nil
}
