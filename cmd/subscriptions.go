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
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription [name] [topic] [amqp-url] [amqp-exchange]",
	Args:  cobra.ExactArgs(4),
	Short: "Create a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()

		name := args[0]
		topic := args[1]
		amqpURL := args[2]

		err := client.RGWCreateSubscription(name, topic, amqpURL)
		if err != nil {
			glog.Fatalf("failed to create subscription: %v", err)
		}
	},
}
