package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version     string = "v1.0.0"
	displayName string = "MySQL Status Watcher"
	VersionCmd         = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of " + displayName,
		Long:  "Print the version number of " + displayName,
		Run: func(cmd *cobra.Command, args []string) {
			// cmd.Printf("%s version %s\n", displayName, Version)
			ShowBanner()
		}}
)

func ShowBanner() {
	banner := "%s %s\n"
	fmt.Printf(banner, displayName, Version)
}
