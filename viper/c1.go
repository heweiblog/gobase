package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	//"reflect"
)

func print_json(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}

func main() {
	viper.SetConfigFile("config.ini") // 指定配置文件
	viper.AddConfigPath(".")          // 指定查找配置文件的路径
	err := viper.ReadInConfig()       // 读取配置信息
	if err != nil {                   // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		all := viper.AllSettings()
		c.JSONP(http.StatusOK, all)
	})
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		fmt.Println(viper.IsSet(name))
		all := viper.Get(name)
		c.JSONP(http.StatusOK, all)
	})

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
