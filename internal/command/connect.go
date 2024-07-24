package command

import (
	"MySQLStatusWatcher/internal/config"
	"MySQLStatusWatcher/internal/connect"
	"MySQLStatusWatcher/internal/logger"

	"github.com/spf13/cobra"
)

var (
	ConnectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Connect to the MySQL database",
		Run: func(cmd *cobra.Command, args []string) {
			if connect_test() {
				logger.Info("Connected to the MySQL database")
			} else {
				logger.Error("Failed to connect to the MySQL database")
			}
		},
	}
)

func connect_test() bool {
	// new connector
	connector := connect.NewConnector("mysql")
	parameters := map[string]interface{}{
		"type": config.Config.MySQL.Type,
		"host": config.Config.MySQL.Host,
		"port": config.Config.MySQL.Port,
		"user": config.Config.MySQL.User,
		"pass": config.Config.MySQL.Pass,
		"db":   config.Config.MySQL.DB,
	}
	connector.Init(parameters)

	flag, err := connector.Connect()
	if err != nil {
		logger.Errorf("type: %s, host: %s, port: %d, user: %s, pass: %s, db: %s", config.Config.MySQL.Type, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.User, config.Config.MySQL.Pass, config.Config.MySQL.DB)
		logger.Errorf("connect error: %s", err.Error())
		return false
	}
	if !flag {
		logger.Errorf("type: %s, host: %s, port: %d, user: %s, pass: %s, db: %s", config.Config.MySQL.Type, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.User, config.Config.MySQL.Pass, config.Config.MySQL.DB)
		logger.Errorf("connect error: %s", err.Error())
		return false
	}
	return true
}
