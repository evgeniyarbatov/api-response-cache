package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	BaseUrl      string `json:"BaseUrl"`
	Tps          int    `json:"Tps"`
	RequestCount int    `json:"RequestCount"`
	UserCount    int    `json:"UserCount"`
}

func (c *Config) ReadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}
