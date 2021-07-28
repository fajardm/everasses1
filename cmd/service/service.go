package service

import (
	"time"

	"github.com/everasses1/config"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

var command = &cobra.Command{
	Use:     "service",
	Aliases: []string{"svc"},
	Short:   "Run service",
	Run: func(c *cobra.Command, args []string) {
		conf, err := config.LoadFile(configFile)
		if err != nil {
			panic(err)
		}

		loc, err := time.LoadLocation(conf.App.Timezone)
		if err != nil {
			panic(err)
		}
		time.Local = loc
	},
}

func init() {
	command.Flags().StringVar(&configFile, "config", "./config/files/config.example.yaml", "Set config file path")
}

// GetCommand .
func GetCommand() *cobra.Command {
	return command
}
