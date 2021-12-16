package cache

import "time"

type CacheExpireTime time.Duration

const (
	CacheExpireTime_permanent CacheExpireTime = 0

	CacheExpireTime_1s  CacheExpireTime = 1
	CacheExpireTime_2s  CacheExpireTime = 2
	CacheExpireTime_3s  CacheExpireTime = 3
	CacheExpireTime_5s  CacheExpireTime = 5
	CacheExpireTime_10s CacheExpireTime = 10
	CacheExpireTime_20s CacheExpireTime = 20
	CacheExpireTime_30s CacheExpireTime = 30

	CacheExpireTime_1min  CacheExpireTime = 60
	CacheExpireTime_5min  CacheExpireTime = 300
	CacheExpireTime_10min CacheExpireTime = 600
	CacheExpireTime_30min CacheExpireTime = 1800

	CacheExpireTime_1h  CacheExpireTime = 3600
	CacheExpireTime_6h  CacheExpireTime = 21600
	CacheExpireTime_12h CacheExpireTime = 43200

	CacheExpireTime_1d  CacheExpireTime = 86400
	CacheExpireTime_3d  CacheExpireTime = 259200
	CacheExpireTime_7d  CacheExpireTime = 604800
	CacheExpireTime_30d CacheExpireTime = 2592000
)
