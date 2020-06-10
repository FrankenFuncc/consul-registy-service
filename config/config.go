package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Logs struct {
	LogFilePath    string `yaml:"LogFilePath"`
	LogLevel       string `yaml:"LogLevel"`
}

type Config struct {
	SiteLogs Logs `yaml:"Logs"`
	System System `yaml:"System"`
}

type System struct {
	ServiceName    string `yaml:"ServiceName"`
}

func (GC *Config) LogConfig() (*Config, error) {
	config, err := ioutil.ReadFile("G:\\goworkbench\\goprojects\\src\\exec\\conf\\consul.yaml")
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
	}
	err = yaml.Unmarshal(config, &GC)
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
	}
	return GC, err
}