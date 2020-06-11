package config

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var confpath string

func init() {
	flag.StringVar(&confpath, "confpath", "conf/consul.yaml", "默认为conf/consul.yaml")
}

type Logs struct {
	LogFilePath    string `yaml:"LogFilePath"`
	LogLevel       string `yaml:"LogLevel"`
}

type Consul struct {
	Token string `yaml:"Token"`
	Address string `yaml:"Address"`
	CheckTimeout string `yaml:"CheckTimeout"`
	CheckInterval string `yaml:"CheckInterval"`
	CheckDeregisterCriticalServiceAfter bool `yaml:"CheckDeregisterCriticalServiceAfter"`
	CheckDeregisterCriticalServiceAfterTime string `yaml:"CheckDeregisterCriticalServiceAfterTime"`
}

type Config struct {
	SiteLogs Logs `yaml:"Logs"`
	System System `yaml:"System"`
	Consul Consul `yaml:"Consul"`
	Service Service `yaml:"Service"`
}

type System struct {
	ServiceName    string `yaml:"ServiceName"`
	ListenAddress  string `yaml:"ListenAddress"`
	Port           string `yaml:"Port"`
	FindAddress    string `yaml:"FindAddress"`
	ConfPath       string `yaml:"ConfPath"`
}

type Service struct {
	Tag    string `yaml:"Tag"`
	Port           string `yaml:"Port"`
	Address string `yaml:"Address"`
}

func (GC *Config) GetConfig() (*Config, error) {
	flag.Parse()
	config, err := ioutil.ReadFile(confpath)
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