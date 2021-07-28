package cmd

import (
	"github.com/everasses1/cmd/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "go-clean-arch",
}

func init() {
	rootCmd.AddCommand(service.GetCommand())
}

// Execute .
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
