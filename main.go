package main

import (
	"exec/config"
	conf "exec/config"
	consul "exec/cons"
	logs "exec/logs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetConf() *conf.Config {
	var conf *conf.Config
	config2, err := conf.GetConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Fatal(err.Error())
	}
	return config2
}
func main() {
	var conf *config.Config
	config, err := conf.GetConfig()
	logs.InitLog(config.SiteLogs.LogFilePath)
	RegistyStart := new(consul.Addresses)
	RegistyStart.GetValues()
	_, err = RegistyStart.CheckSorted("node-exporter")
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
	}
	err = RegistyStart.CheckAddr("node-exporter")
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
		panic("Error" + err.Error())
	}
	//定义一个http接口
	http.HandleFunc("/", consul.Handler)
	err = http.ListenAndServe(GetConf().System.ListenAddress + ":" + GetConf().System.Port, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
		panic("Error" + err.Error())
	}
}
