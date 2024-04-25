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
	Db     db     `yaml:"db"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
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
