package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
    "strconv"
)

var tubeKickCommand = &cobra.Command{
    Use:   "kick [tube_name] [num]",
    Short: "Kick [num] jobs into the ready queue",
    Long:  `Everyone loves soccer.`,
}

func init() {
    tubeKickCommand.Run = tube_kick
}

func tube_kick(cmd *cobra.Command, args []string) {
    if len(args) < 2 {
        cmd.Help()
        return
    }

    num64, err := strconv.ParseUint(args[1], 10, 0)
    if err != nil {
        log.Fatal(err)
    }

    num := int(num64)

    client := connect()
    defer client.Close()

    client.Tube = beanstalk.Tube{client, tube_name}

    kicked, err := client.Tube.Kick(num)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Kicked: [%d] jobs.\n", kicked)
}
