package tnet

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"tinynet/tutil"
)

type Config struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	MaxConn       int    `json:"maxConn"`        //最大的连接数
	MaxPacketSize int    `json:"maxPacketSize"`  //最大的数据包
	ConfFilePath  string `json:"configFilePath"` //配置文件url
	MaxMsgChanLen int    `json:"maxMsgChanLen"`  //最大msg缓冲
}

var ConfigObj *Config

func Init() {
	ConfigObj = &Config{
		Host:          "0.0.0.0",
		Port:          9999,
		Name:          "tinyServer",
		Version:       "v1.0",
		MaxConn:       100,
		MaxPacketSize: 4096,
		ConfFilePath:  "config/config.json",
		MaxMsgChanLen: 100,
	}
	ConfigObj.LoadConfig()
}

func (c *Config) LoadConfig() {
	if confFileExists, _ := tutil.PathExists(c.ConfFilePath); confFileExists != true {
		panic(errors.New("config file not exists "))
	}

	data, err := ioutil.ReadFile(c.ConfFilePath)
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	//fmt.Printf("json :%s\n", data)
	err = json.Unmarshal(data, &ConfigObj)
	if err != nil {
		panic(err)
	}
}
