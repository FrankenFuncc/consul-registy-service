package common

import (
	conf "exec/config"
	"github.com/sirupsen/logrus"
)

func GetConf() (*conf.Config, error) {
	var config conf.Config
	config2, err := config.GetConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Fatal(err.Error())
	}
	return config2, err
}
