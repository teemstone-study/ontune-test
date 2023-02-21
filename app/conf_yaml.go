package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigDbInfo struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ConfigScrape struct {
	Rate int `yaml:"rate"`
	Avg  int `yaml:"avg"`
}

type ConfigHost struct {
	IgnoreProcessorCountChange int          `yaml:"ignoreprocessorcountchange"`
	Active                     int          `yaml:"active"`
	Perf                       ConfigScrape `yaml:"perf"`
	Proc                       ConfigScrape `yaml:"proc"`
	Disk                       ConfigScrape `yaml:"disk"`
	Net                        ConfigScrape `yaml:"net"`
	CPU                        ConfigScrape `yaml:"cpu"`
}

type ConfigDemo struct {
	HostCount       int `yaml:"host_count"`
	HostChangeCount int `yaml:"host_change_count"`
	ChangeInterval  int `yaml:"change_interval"`
}

type Config struct {
	Database []ConfigDbInfo `yaml:"database"`
	Host     ConfigHost     `yaml:"host"`
	Demo     ConfigDemo     `yaml:"demo"`
}

func (d ConfigDbInfo) Datasource() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul",
		d.Host, d.Port, d.Username, d.Password, d.Dbname)
}

func GetConfig(filename string) Config {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
