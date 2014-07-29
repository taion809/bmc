package commands

import (
    "github.com/spf13/cobra"
)

var initCommand = &cobra.Command{Use: "bmc"}

var hostname, port, protocol string

func AddCommands() {
    initCommand.AddCommand(listTubesCmd)
    initCommand.AddCommand(debugCmd)
    initCommand.AddCommand(tubeCommand)
    tubeCommand.AddCommand(tubeStatsCommand)
    tubeCommand.AddCommand(tubePeekReadyCommand)
    tubeCommand.AddCommand(tubePeekBuriedCommand)
    tubeCommand.AddCommand(tubePeekDelayedCommand)
    tubeCommand.AddCommand(tubeKickCommand)
}

func Execute() {
    AddCommands()
    initCommand.Execute()
}

func init() {
    initCommand.PersistentFlags().StringVarP(&hostname, "hostname", "h", "localhost", "beanstalkd hostname or IPv4 address")
    initCommand.PersistentFlags().StringVar(&protocol, "protocol", "tcp", "transport protocol")
    initCommand.PersistentFlags().StringVarP(&port, "port", "p", "11300", "listening port")
}
