package commands

import (
	"github.com/spf13/cobra"
)

// This is meant to be used as an anchor for other tube based subcommands
var tubeCommand = &cobra.Command{
	Use:   "tube",
	Short: "Tube based subcommands",
	Long:  `Enjoy working with the beanstalkd tubes, this command has it all!`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	tubeCommand.PersistentFlags().StringVarP(&tube_name, "tube_name", "t", "default", "tube to operate on")
}
