package cache

type RedisResult struct {
	err error
	val interface{}
}

func (result *RedisResult) OK() bool {
	return result.err == nil
}

func (result *RedisResult) Error() error {
	return result.err
}

func (result *RedisResult) StringVal() string {
	return result.val.(string)
}

func (result *RedisResult) BoolVal() bool {
	return result.val.(bool)
}

func (result *RedisResult) Int64Val() int64 {
	return result.val.(int64)
}

func (result *RedisResult) Float64Val() float64 {
	return result.val.(float64)
}

func (result *RedisResult) StringArrayVal() []string {
	return result.val.([]string)
}

func (result *RedisResult) StringMapVal() map[string]string {
	return result.val.(map[string]string)
}
