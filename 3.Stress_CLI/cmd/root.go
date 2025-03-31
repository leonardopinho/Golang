package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress_app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		//ctx := cmd.Context()
		url, _ := cmd.Flags().GetString("url")
		fmt.Println(url)
		//totalRequestCount, _ := cmd.Flags().GetInt("requests")
		//concurrency, _ := cmd.Flags().GetInt("concurrency")

		//requester.Request(ctx, url, concurrency, totalRequestCount)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("url", "", "set target URL")
	rootCmd.MarkPersistentFlagRequired("url")
}
