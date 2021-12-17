package cache

type UserCacheKeys struct {
	*CacheKeys
}

var UserKeys = &UserCacheKeys{
	&CacheKeys{
		prefix: "user",
	},
}

func (keys *UserCacheKeys) Detail(mid int64) string {
	return keys.format("detail:%d", mid)
}

func (keys *UserCacheKeys) AddressList(mid int64) string {
	return keys.format("address_list:%d", mid)
}
