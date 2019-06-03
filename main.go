package main

import (
	"parse_conf/cache"
	cfg "parse_conf/config"	
  "fmt"
)


func main() {
	// 加载配置信息
	instance := cfg.GetAppInstance()
  info = fmt.Sprintf("#####  configuration:\n"+
		"[redis] 		    server   			: %s\n"+
		"[redis] 		    password 			: %s\n"+
		"[redis] 		    db       			: %d\n",
		conf.GetRedisHosts(), conf.GetRedisPassword(),
		conf.GetRedisDb())
	fmt.Println(info)
	cache.CacheDataSource().DataSourceInit(instance.GetRedisHosts(),
		instance.GetRedisPassword(), instance.GetRedisDb())
}
