package commands

import (
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var initCommand = &cobra.Command{Use: "bmc"}

var hostname, port, protocol string

func AddCommands() {
    initCommand.AddCommand(debugCmd)
    initCommand.AddCommand(listTubesCmd)
    initCommand.AddCommand(buryCommand)
    initCommand.AddCommand(deleteCommand)
    initCommand.AddCommand(peekCommand)
    initCommand.AddCommand(connStatsCommand)
    initCommand.AddCommand(jobStatsCommand)
    initCommand.AddCommand(touchCommand)
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
    initCommand.PersistentFlags().StringVarP(&port, "port", "p", "11301", "listening port")
}

func connect() *beanstalk.Conn {
    client, err := beanstalk.Dial(protocol, hostname+":"+port)
    if err != nil {
        log.Fatal(err)
    }

    return client
}
