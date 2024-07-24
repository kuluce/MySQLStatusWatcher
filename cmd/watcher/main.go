package main

import (
	"MySQLStatusWatcher/internal/command"
	"MySQLStatusWatcher/internal/config"
	"MySQLStatusWatcher/internal/logger"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "A simple MySQL status watcher",
		Long:  "A simple MySQL status watcher",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	// code here
		// },
		Version: "v1.0.0",
	}
)

func main() {
	logger.InitLog("watcher")

	if err := config.LoadConfig(); err != nil {
		logger.Errorf("load config error: %s", err.Error())
		return
	}

	rootCmd.AddCommand(command.VersionCmd)
	rootCmd.AddCommand(command.EncryptCmd)
	rootCmd.AddCommand(command.ConnectCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
