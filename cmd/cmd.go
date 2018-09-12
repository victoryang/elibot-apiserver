package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var rootCmd := &cobra.Command {
		Use:	"elibot-server [option]",
		Short:	"An api server for elibot",
		Long:	`An api server for elibot written in go, 
				providing a generic api service to control and get data from robot`,
		Run:	runDaemon,
	}

	flags := rootCmd.Flags()
	flags.PersistentFlags().StringVarP(&configFile, "configure", "C", "/rbctrl/Config/elibot-server.yaml", 
																		"configuration file for api server")
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}