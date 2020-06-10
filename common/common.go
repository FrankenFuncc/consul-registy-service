package common

import (
	conf "exec/config"
	"log"
)

func GetLogConf() (*conf.Config, error) {
	var config conf.Config
	config2, err := config.LogConfig()
	if err != nil {
		log.Println("ERR>>>", err.Error())
	}
	return config2, err
}
