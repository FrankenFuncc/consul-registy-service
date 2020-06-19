package main

import (
	conf "exec/config"
	consul "exec/cons"
	logs "exec/logs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logs.InitLog(conf.GetConf().SiteLogs.LogFilePath)
	RegistyStart := new(consul.Addresses)
	_, err := RegistyStart.CheckSorted(conf.GetConf().Service.Tag)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
	}
	err = RegistyStart.CheckAddr(conf.GetConf().Service.Tag)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
		panic("Error" + err.Error())
	}
	//定义一个http接口
	http.HandleFunc("/", consul.Handler)
	err = http.ListenAndServe(conf.GetConf().System.ListenAddress+":"+conf.GetConf().System.Port, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Info(err.Error())
		panic("Error" + err.Error())
	}
}
