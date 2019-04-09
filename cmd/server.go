package cmd

import (
	"github.com/gmendonca/roboat/pkg/webserver"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server command line interface",
	Long:  `server, the command line interface`,
	Run: func(cmd *cobra.Command, args []string) {
		w := &webserver.Webserver{
			Port: ":3000",
		}
		w.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
