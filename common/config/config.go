package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  config.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 14:42
 */

// 文件配置
type config struct {
	Server server `yaml:"server"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// Config 全局变量
var Config *config

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	// 将yaml文件中的内容解析到Config变量中
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
