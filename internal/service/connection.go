package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return []*define.Connection{{
			Identity: "1",
			Name:     "1",
			Addr:     "1",
			Port:     "1",
			Username: "1",
			Password: "1",
		}}, nil
	}
	conf := new(define.Config)
	json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}
