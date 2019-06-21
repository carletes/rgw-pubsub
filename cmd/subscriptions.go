package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	ackLevel string
)

func init() {
	createSubscriptionCmd.Flags().StringVar(&ackLevel, "amqp-ack-level", "none", "AMQP ack level")
	rootCmd.AddCommand(createSubscriptionCmd)
	rootCmd.AddCommand(deleteSubscriptionCmd)
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription [name] [topic] [amqp-url] [amqp-exchange]",
	Args:  cobra.ExactArgs(4),
	Short: "Create a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()

		name := args[0]
		topic := args[1]
		url := args[2]
		exchange := args[3]

		err := client.RGWCreateSubscription(name, topic, url, exchange, ackLevel)
		if err != nil {
			glog.Fatalf("failed to create subscription: %v", err)
		}
	},
}

var deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete-subscription",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()

		name := args[0]
		err := client.RGWDeleteSubscription(name)
		if err != nil {
			glog.Fatalf("failed to delere subscription %s: %v", name, err)
		}
	},
}
