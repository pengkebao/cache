package cache

func init() {
	//初始化成redis
	opts := &RedisOpts{
		Host: "127.0.0.1:6379",
	}
	Instance = NewRedis(opts)
	//初始化成memcache
	//Instance = NewMemcache("127.0.0.1:11211")

	//初始化成内存缓存
	// defaultExpiration := time.Hour // The default for the default is one hour.
	// Instance = NewInMemoryCache(defaultExpiration)
}
