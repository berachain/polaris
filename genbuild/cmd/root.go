package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "genbuild",
}

func init() {
	initCmd, err := NewInitCmd()
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(initCmd)

	collectCmd, err := NewCollectCmd()
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(collectCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
