package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

/*
 * @Title readconfig.go
 * @Description 读取配置文件
 * @Author  zongweizheng 2022/8/22 16:45
 * @Update
 */

type Conf struct {
	Server struct {
		Debug string `yaml:"debug"`
	}
	MYSQL struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	LOG struct {
		Level string `yaml:"level"`
	}
}

func GetConfig(path string) Conf {
	var conf Conf
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	return conf
}
