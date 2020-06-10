package main

import (
	"exec/config"
	consul "exec/cons"
	logs "exec/logs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	var conf *config.Config
	config, err := conf.LogConfig()
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
	err = http.ListenAndServe("0.0.0.0:9983", nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
		panic("Error" + err.Error())
	}
}
