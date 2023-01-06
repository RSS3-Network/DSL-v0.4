package databeat

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

const (
	PREFIX = "[DATABEAT]"
)

// Model struct
type DataModel struct {
	Name   string
	Fields []string
}

var (
	beatLogger *log.Logger
	once       sync.Once
)

// Singleton beat logger
func GetBeatLogger() *log.Logger {
	once.Do(func() {
		beatLogger = log.New(os.Stderr, PREFIX, 0)
	})
	return beatLogger
}

// Log to stderr
func (d *DataModel) Beat(content map[string]interface{}) (string, error) {
	err := d.validate(content)
	if err != nil {
		return "", err
	}

	content["index"] = d.Name

	dataBytes, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	dataStr := string(dataBytes)

	GetBeatLogger().Println(dataStr)
	return dataStr, nil
}

// Validate DataModel and content
func (d *DataModel) validate(content map[string]interface{}) error {
	// validate DataModel, panic
	if d.Name == "" {
		panic("缺少name")
	}
	if d.Fields == nil {
		panic("缺少fields")
	}
	for _, field := range d.Fields {
		if field[0] == '@' {
			panic("不能以@开始")
		}
	}

	// validate content, return err
	for _, field := range d.Fields {
		if _, ok := content[field]; !ok {
			return errors.New("缺少字段: " + field)
		}
	}
	if len(content) != len(d.Fields) {
		return errors.New("传入的内容字段不合法")
	}

	return nil
}
