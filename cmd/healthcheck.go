package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Check healthcheck",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ok")
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}
