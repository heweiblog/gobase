package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"reflect"
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
	//viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.SetConfigFile("config.ini") // 指定配置文件
	viper.AddConfigPath(".")          // 指定查找配置文件的路径
	err := viper.ReadInConfig()       // 读取配置信息
	if err != nil {                   // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.GET("/", func(c *gin.Context) {
		all := viper.AllSettings()
		print_json(all)
		fmt.Println(all)
		b, err := json.Marshal(all)
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}
		fmt.Println("b:", string(b))
		for k, v := range all {
			fmt.Println(k, v, reflect.TypeOf(v))
			//for k1, v1 := range v {
			//fmt.Println("=========", k1, v1, reflect.TypeOf(v1))
			//}
		}
		//c.JSON(http.StatusOK, gin.H{"status": "OK"})
		c.JSON(http.StatusOK, string(b))
	})

	if err := r.Run(
		//fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		fmt.Sprintf(":%d", 8000)); err != nil {
		panic(err)
	}
}
