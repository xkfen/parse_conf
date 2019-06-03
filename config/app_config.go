package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"github.com/widuu/goini"
)

const DefaultConfigFile = "/etc/license.conf"

type appConfig struct {
	redisHosts    []string
	redisPassword string
	redisDb       int
}

var appInstance *appConfig
var appOnce sync.Once

func GetAppInstance() *appConfig {
	appOnce.Do(func() {
		appInstance = &appConfig{}
		pwd, _ := os.Getwd()
		confFile := pwd + DefaultConfigFile
		conf := goini.SetConfig(confFile)
		redisHosts := conf.GetValue("redis", "host")
		redisPassword := conf.GetValue("redis", "password")
		redisDb := conf.GetValue("redis", "db")
		appInstance.redisHosts = strings.Split(redisHosts, ",")
		if len(redisPassword) == 0 {
			appInstance.redisPassword = ""
		} else {
			appInstance.redisPassword = redisPassword
		}
		if len(redisDb) == 0 {
			appInstance.redisDb = 0
		} else {
			appInstance.redisDb, _ = strconv.Atoi(redisDb)
		}
	})
	return appInstance
}

func (a *appConfig) GetRedisHosts() []string {
	return a.redisHosts
}

func (a *appConfig) GetRedisPassword() string  {
	return a.redisPassword
}

func (a *appConfig) GetRedisDb () int {
	return a.redisDb
}
