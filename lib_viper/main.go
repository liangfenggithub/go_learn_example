package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	/*
		viper 的使用非常简单，它需要很少的设置。设置文件名（SetConfigName）、配置类型（SetConfigType）和搜索路径（AddConfigPath），然后调用ReadInConfig。
		viper会自动根据类型来读取配置。使用时调用viper.Get方法获取键值。

				有几点需要注意：

		设置文件名时不要带后缀；
		搜索路径可以设置多个，viper 会根据设置顺序依次查找；
		viper 获取值时使用section.key的形式，即传入嵌套的键名；
		默认值可以调用viper.SetDefault设置。
	*/

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	/*Get方法返回一个interface{}的值，使用有所不便。
	GetType系列方法可以返回指定类型的值。
	其中，Type 可以为Bool/Float64/Int/String/Time/Duration/IntSlice/StringSlice。
	但是请注意，如果指定的键不存在或类型不正确，GetType方法返回对应类型的零值。

	如果要判断某个键是否存在，使用IsSet方法。
	另外，GetStringMap和GetStringMapString直接以 map 返回某个键下面所有的键值对，前者返回map[string]interface{}，后者返回map[string]string。
	AllSettings以map[string]interface{}返回所有设置。


	*/
	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))
}
