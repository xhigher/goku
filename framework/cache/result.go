package cache

import "goku.net/utils"

type RedisResult struct {
	err   error
	exist bool
	val   interface{}
}

func (result *RedisResult) OK() bool {
	return result.err == nil
}

func (result *RedisResult) Exist() bool {
	return result.exist
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

func (result *RedisResult) String2Struct(data interface{}) error {
	return utils.ParseStruct(result.StringVal(), data)
}
