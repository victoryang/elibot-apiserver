package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command {
		Use:	"elibot-server [option]",
		Short:	"An api server for elibot",
		Long:	`An api server for elibot written in go, providing a generic api service to control and get data from robot`,
		RunE:	runDaemon,
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&configFile, "configure", "C", "/rbctrl/Config/elibot-server.yaml", "configuration file for api server")
	return rootCmd
}


func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
