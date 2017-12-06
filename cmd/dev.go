package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"aews/core/model/api"
	"aews/core/service"
)

func init() {
        RootCmd.AddCommand(devCmd)
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start Dev",
	Long:  `Start Dev`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
			case "alerter":
				service.AlerterService().SnedDayAlerter()
			case "SlackApi":
				api.SlackApi().SendMessage("數據異常警示", "單一遊戲場次累積10萬場，RTP > 100%現象")
			default:
				fmt.Print("default")
			}
	},
}
