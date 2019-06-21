package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createTopicCmd)
	rootCmd.AddCommand(deleteTopicCmd)
}

var createTopicCmd = &cobra.Command{
	Use:   "create-topic [name]",
	Short: "Create a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		topic := args[0]
		err := client.RGWCreateTopic(topic)
		if err != nil {
			glog.Fatalf("failed to create topic %s: %v", topic, err)
		}
	},
}

var deleteTopicCmd = &cobra.Command{
	Use:   "delete-topic [name]",
	Short: "Delete a topic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		topic := args[0]
		err := client.RGWDeleteTopic(topic)
		if err != nil {
			glog.Fatalf("failed to delete topic %s: %v", topic, err)
		}
	},
}
