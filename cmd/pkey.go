package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/werbot/lime/license"
)

var pairKeyCmd = &cobra.Command{
	Use:   "pkey",
	Short: "Generating key pair",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		keyPair := license.KeyPairGenerate()
		b, err := json.Marshal(keyPair)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(pairKeyCmd)
}
