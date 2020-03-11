package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

func NewConfig() Config {
	var c Config
	fileName := "../config.yaml"
	_, _fileName, _, _ := runtime.Caller(1)              // get the caller path
	filePath := path.Join(path.Dir(_fileName), fileName) //Dir() split the elem behind of "/"
	f, err := ioutil.ReadFile(filePath)                  //读取yaml配置文件
	if err != nil {
		log.Fatal(err)
	}

	// 解析配置文件
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	return c
}
