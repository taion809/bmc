package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var tubeKickCommand = &cobra.Command{
    Use:   "kick [tube_name] [# num]",
    Short: "Kick [num] jobs into the ready queue",
    Long:  `Everyone loves soccer.`,
}

func init() {
    tubeKickCommand.Run = tube_kick
}

func tube_kick(cmd *cobra.Command, args []string) {
    client, err := beanstalk.Dial(protocol, hostname+":"+port)
    if err != nil {
        log.Fatal(err)
    }

    defer client.Close()

    client.Tube = beanstalk.Tube{client, tube_name}

    id, body, err := client.Tube.PeekBuried()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Delayed: [%d] : %s\n", id, body)
}
