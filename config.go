package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)
var QPS = 1000
var IP = []string{"192.168.1.101"}
//设置配置中心
func GetConsualConfig(host string,port int64,prefix string)(config.Config,error){
	consualSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host + ":" + strconv.FormatInt(port, 10)),
		//设置前缀
		consul.WithPrefix(prefix),
		//设置是否移除前缀,表示可以不带前缀直接获取对应的地址
		consul.StripPrefix(true),
		)
	//配置初始化`
	newConfig,err := config.NewConfig()
	if err != nil{
		return newConfig,err
	}
	//加载配置
	err = newConfig.Load(consualSource)
	return newConfig,err
}