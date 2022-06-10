package service

import (
	"io/ioutil"

	"github.com/talbx/celepush/internal/pkg/model"
	"gopkg.in/yaml.v3"
)

type ConfigReader interface {
	ReadConfig(config *model.AppConfig)
}

type ConfigReaderImpl struct{}

var AppConf model.AppConfig

func (reader ConfigReaderImpl) ReadConfig(config *model.AppConfig) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, config)
	AppConf = *config
	if err != nil {
		panic(err)
	}
}
