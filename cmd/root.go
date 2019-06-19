package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	AccessKey      string
	SecretKey      string
	PubSubEndpoint string
	Username       string
	Zonegroup      string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&AccessKey, "access-key", "", "S3 access key")
	rootCmd.PersistentFlags().StringVar(&PubSubEndpoint, "pubsub-endpoint", "", "URL of PubSub gateway")
	rootCmd.PersistentFlags().StringVar(&SecretKey, "secret-key", "", "S3 secret key")
	rootCmd.PersistentFlags().StringVar(&Username, "user-name", "", "S3 user name")
	rootCmd.PersistentFlags().StringVar(&Zonegroup, "zone-group", "", "RGW zone group")
}

var rootCmd = &cobra.Command{
	Use:  "rgw-pubsub",
	Long: "A command-line tool to interact with Ceph's object store PubSub system.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
