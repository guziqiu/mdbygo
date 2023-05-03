package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	conf := new(define.Config)
	json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

// connectCreate 创建连接
func ConnectionCreate(conn *define.Connection) error {
	if conn.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	// 参数默认值处理
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}

	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	confName := nowPath + string(os.PathSeparator) + define.ConfigName
	data, err := ioutil.ReadFile(confName)
	if errors.Is(err, os.ErrNotExist) {
		// 配置文件内容初始化
		conf.Connections = []*define.Connection{conn}
		data, _ := json.Marshal(conf)
		// 写入配置文件
		os.MkdirAll(nowPath, 0666)
		ioutil.WriteFile(confName, data, 0666)
	}
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, conn)
	data, error := json.Marshal(conf)
	if nil != error {
		fmt.Println(error)
	}
	ioutil.WriteFile(confName, data, 0666)

	return nil
}
