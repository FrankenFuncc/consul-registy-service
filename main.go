package main

import (
	conf "exec/config"
	consul "exec/cons"
	logs "exec/logs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetConf() *conf.Config {
	var conf2 *conf.Config
	config2, err := conf2.GetConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Fatal(err.Error())
	}
	return config2
}
func main() {
	logs.InitLog(conf.GetConf().SiteLogs.LogFilePath)
	RegistyStart := new(consul.Addresses)
	_, err := RegistyStart.CheckSorted("node-exporter")
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
	}
	err = RegistyStart.CheckAddr("node-exporter")
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
		panic("Error" + err.Error())
	}
	//定义一个http接口
	http.HandleFunc("/", consul.Handler)
	err = http.ListenAndServe(GetConf().System.ListenAddress+":"+GetConf().System.Port, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
		panic("Error" + err.Error())
	}
}
