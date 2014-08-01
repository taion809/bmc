package commands

import (
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var initCommand = &cobra.Command{Use: "bmc"}

var (
    ttr       int64
    delay     int64
    pri       uint32
    filename  string
    tube_name string
    hostname  string
    port      string
    protocol  string
)

func AddCommands() {
    initCommand.AddCommand(debugCmd)
    initCommand.AddCommand(buryCommand)
    initCommand.AddCommand(deleteCommand)
    initCommand.AddCommand(listTubesCmd)
    initCommand.AddCommand(peekCommand)
    initCommand.AddCommand(connStatsCommand)
    initCommand.AddCommand(jobStatsCommand)
    initCommand.AddCommand(touchCommand)
    initCommand.AddCommand(releaseCommand)
    initCommand.AddCommand(tubeCommand)
    tubeCommand.AddCommand(tubePutCommand)
    tubeCommand.AddCommand(tubeStatsCommand)
    tubeCommand.AddCommand(tubePeekReadyCommand)
    tubeCommand.AddCommand(tubePeekBuriedCommand)
    tubeCommand.AddCommand(tubePeekDelayedCommand)
    tubeCommand.AddCommand(tubeKickCommand)
    tubeCommand.AddCommand(tubePauseCommand)
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
