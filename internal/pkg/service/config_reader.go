package service

import (
	"fmt"
	"github.com/talbx/birthday-notice/internal/pkg/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type ConfigReader interface {
	ReadConfig() model.AppConfig
}

type ConfigReaderImpl struct {}

var AppConf, _ = ConfigReaderImpl{}.ReadConfig()

func (reader ConfigReaderImpl) ReadConfig() (*model.AppConfig, error) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &model.AppConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}

func ReadCandidates(entries *[]model.BirthdayEntry) {
	filename := "bdays.yaml"
	yamlFile, err := ioutil.ReadFile(filename)

	_ = fmt.Sprintf("---\n%s\n\n", string(yamlFile))
	// log.Println(sprintf)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, entries)
	if err != nil {
		log.Fatal(err)
	}
}
