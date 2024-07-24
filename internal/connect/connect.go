package connect

import (
	"MySQLStatusWatcher/internal/aesencrypt"
	"MySQLStatusWatcher/internal/logger"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connector interface {
	Connect() (bool, error)
	Init(parameters map[string]interface{})
}

type MySQL struct {
	Type string `json:"type" toml:"type" yaml:"type"`
	Host string `json:"host" toml:"host" yaml:"host"`
	Port int    `json:"port" toml:"port" yaml:"port"`
	User string `json:"user" toml:"user" yaml:"user"`
	Pass string `json:"pass" toml:"pass" yaml:"pass"`
	DB   string `json:"db" toml:"db" yaml:"db"`
}

func (m *MySQL) Connect() (bool, error) {
	password := aesencrypt.Decrypt(m.Pass)

	var dsn string
	if m.Type == "socket" {
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s", m.User, password, m.Host, m.DB)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.User, password, m.Host, m.Port, m.DB)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// show status
	rows, err := db.Query("show status")
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var Variable_name string
		var Value string
		err = rows.Scan(&Variable_name, &Value)
		if err != nil {
			logger.Errorf("scan mysql error: %v", err)
			return false, err
		}
		// logger.Infof("Variable_name: %s, Value: %s", Variable_name, Value)
	}

	return true, nil
}

func (m *MySQL) Init(parameters map[string]interface{}) {
	m.Type = parameters["type"].(string)
	m.Host = parameters["host"].(string)
	m.User = parameters["user"].(string)
	m.Pass = parameters["pass"].(string)
	m.DB = parameters["db"].(string)
	m.Port = parameters["port"].(int)
}

func NewConnector(db string) Connector {
	switch db {
	case "mysql":
		return &MySQL{}
	default:
		return nil
	}
}
