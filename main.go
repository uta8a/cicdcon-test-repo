package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "cicdcon-cli",
	Long: "サンプルテストアプリだよ",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := cmd.Flags().GetBool("flag")
		if err != nil {
			return err
		}

		fmt.Printf("Flag is %t", t)
		return nil
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("flag", "f", false, "サンプルフラグ")
}
