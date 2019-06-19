package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createNotificationCmd)
}

var createNotificationCmd = &cobra.Command{
	Use:   "create-notification [bucket] [topic]",
	Args:  cobra.ExactArgs(2),
	Short: "Create a notification",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()

		bucket := args[0]
		topic := args[1]
		err := client.RGWCreateNotification(bucket, topic)
		if err != nil {
			glog.Fatalf("failed to create notification for bucket %s and topic %s: %v", bucket, topic, err)
		}
	},
}
