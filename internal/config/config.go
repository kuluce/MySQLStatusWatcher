package config

import (
	"MySQLStatusWatcher/internal/logger"
	"os"
	"path/filepath"

	"github.com/kuluce/pkg/file"
	"gopkg.in/yaml.v3"
)

type mysqlConf struct {
	Type string `json:"type" toml:"type" yaml:"type"`
	Host string `json:"host" toml:"host" yaml:"host"`
	Port int    `json:"port" toml:"port" yaml:"port"`
	User string `json:"user" toml:"user" yaml:"user"`
	Pass string `json:"pass" toml:"pass" yaml:"pass"`
	DB   string `json:"db" toml:"db" yaml:"db"`
}

type conf struct {
	MySQL mysqlConf `json:"mysql" toml:"mysql" yaml:"mysql"`
}

func (c *conf) Save() error {

	_content, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	_configFile := filepath.Join(file.SelfDir(), "config.yaml")
	_, err = file.WriteBytes(_configFile, _content)
	if err != nil {
		return err
	}

	return nil
}

var Config *conf

func LoadConfig() error {
	_configFile := filepath.Join(file.SelfDir(), "config.yaml")
	if !file.IsExist(_configFile) {
		Config = NewDefaultConfig()
	}

	_content, err := os.ReadFile(_configFile)
	if err != nil {
		return err
	}

	var _conf conf
	err = yaml.Unmarshal(_content, &_conf)
	if err != nil {
		return err
	}

	Config = &_conf

	return nil
}

func NewDefaultConfig() *conf {
	_conf := conf{
		MySQL: mysqlConf{
			Type: "socket",
			Host: "/var/run/mysqld/mysqld.sock",
			Port: 3306,
			User: "root",
			Pass: "root",
			DB:   "mysql",
		},
	}

	if err := _conf.Save(); err != nil {
		logger.Errorf("save default config error: %v", err)
	}

	return &_conf
}
