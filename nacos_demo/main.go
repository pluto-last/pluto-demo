package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func main() {
	sc := []constant.ServerConfig{{
		IpAddr: "127.0.0.1",
		Port:   8848,
	}}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "df28523a-2b30-4ddb-a2e3-6a22cca82af8", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./nacos/log",
		CacheDir:            "./nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		log.Println("CreateConfigClient err", err)
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "pluto_config.yaml",
		Group:  "pluto",
	})

	if err != nil {
		log.Println("get config err", err)
		panic(err)
	}
	fmt.Println(content) //字符串 - yaml

	log.Println("content == ", content)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "pluto_config.yaml",
		Group:  "pluto",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件发生了变化...")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	select {}
}
