package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Mongo struct {
	Uri string `json:"uri"`
}

type Mysql struct {
	Pass string `json:"pass"`
	Uri  string `json:"uri"`
}

type Config struct {
	Env        string `json:"env,omitempty"`
	Port       string `json:"port,omitempty"`
	Mongo      Mongo  `json:"mongo,omitempty"`
	Mysql      Mysql  `json:"mysql,omitempty"`
	DefaultVar string `json:"default_var,omitempty"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var Conf Config

func Init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	v := viper.New()
	configType := "json"

	v.AddConfigPath("./config")
	v.SetConfigName("default")
	v.SetConfigType(configType)

	// log.Println("read config [default] ...")
	if err := v.ReadInConfig(); err != nil {
		panic("get default.config:" + err.Error())
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入，否则无法达到local覆盖default的目的。
	for k, value := range configs {
		v.SetDefault(k, value)
	}

	env := os.Getenv("ENV")

	if env != "" {
		// log.Printf("read config [%s] ...\n", env)
		v.SetConfigName(env)
		err := v.ReadInConfig()
		if err != nil {
			panic("get default.config:" + err.Error())
		}
	} else {
		env = "default"
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("config Unmarshal err:" + err.Error())
	}

	log.Printf("env: %s, config loaded.", env)
}
