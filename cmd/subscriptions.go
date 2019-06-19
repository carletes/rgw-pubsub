package cmd

import (
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
	Args:  cobra.ExactArgs(5),
	Short: "Create a subscription",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
