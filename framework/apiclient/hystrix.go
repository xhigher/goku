package apiclient

import "github.com/afex/hystrix-go/hystrix"

func InitHystrix() {
	config := hystrix.CommandConfig{
		Timeout:                2000, //超时时间设置  单位毫秒
		MaxConcurrentRequests:  500,  //最大请求数
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率
		RequestVolumeThreshold: 5,    //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	}
	hystrix.ConfigureCommand("goku-http", config)
}

var (
	hystrixConfig = hystrix.CommandConfig{
		Timeout:                2000,  //超时时间设置  单位毫秒
		MaxConcurrentRequests:  500,   //最大请求数
		SleepWindow:            50000, //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  10,    //错误率
		RequestVolumeThreshold: 5,     //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	}
)
