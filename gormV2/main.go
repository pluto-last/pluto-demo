package main

import (
	"flag"
	"gV2/config"
	"gV2/global"
	"gV2/middleware/cache"
	"gV2/middleware/db"
	myzap "gV2/middleware/zap"

	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

var pacPort = 4567
var socksPort = 5678
var configPath string

func main() {

	flag.StringVar(&configPath, "config", "./conf/conf.yaml", "config file path")
	flag.Parse()

	// 资源初始化
	InitServer()

	findInMap()
}

func InitServer() {
	var err error
	global.GVA_CONFIG, err = config.Get(configPath) // 初始化配置
	if err != nil {
		log.Fatalln("read conf file : ", err)
	}

	global.GVA_LOG = myzap.Zap() // 初始化zap日志库

	global.GVA_DB, err = db.Instance(global.GVA_CONFIG) // 初始化数据库库
	if err != nil {
		log.Fatalln("init db err : ", err)
	}

	global.GVA_REDIS, err = cache.InitRedis(global.GVA_CONFIG.Cache.Addr, global.GVA_CONFIG.Cache.Password, 0) // 初始化redis服务
	if err != nil {
		log.Fatalln("redis init err : ", err)
	}

}
