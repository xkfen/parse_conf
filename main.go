package main

import (
	"github.com/parse_conf/cache"
	cfg "github.com/parse_conf/config"	
 	 "fmt"
)


func main() {
	// 加载配置信息
	instance := cfg.GetAppInstance()
  info := fmt.Sprintf("#####  configuration:\n"+
		"[redis] 		    server   			: %s\n"+
		"[redis] 		    password 			: %s\n"+
		"[redis] 		    db       			: %d\n",
	  instance.GetRedisHosts(), instance.GetRedisPassword(),
	  instance.GetRedisDb())
	fmt.Println(info)
	cache.CacheDataSource().DataSourceInit(instance.GetRedisHosts(),
		instance.GetRedisPassword(), instance.GetRedisDb())
}
