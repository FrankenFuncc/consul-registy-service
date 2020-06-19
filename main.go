package main

import (
	conf "exec/config"
	consul "exec/cons"
	logs "exec/logs"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {

	//time.Sleep(time.Duration(20) * time.Second)
	logs.InitLog(conf.GetConf().SiteLogs.LogFilePath)
Check:
	if !consul.GetSvcCode() {
		logrus.WithFields(logrus.Fields{}).Info("检测到服务端口未启动，等待启动...")
		time.Sleep(time.Duration(2) * time.Second)
		goto Check
	} else {
		logrus.WithFields(logrus.Fields{}).Info("检测到服务端口启动,开始注册...")
	}
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
