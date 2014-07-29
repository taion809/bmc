package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var tubePeekBuriedCommand = &cobra.Command{
    Use:   "peek_buried [tube_name]",
    Short: "Retrieve the first available job in the buried queue",
    Long:  `Buried jobs, kick it back to life or lay it to rest.`,
}

func init() {
    tubePeekBuriedCommand.Run = tube_peek_buried
}

func tube_peek_buried(cmd *cobra.Command, args []string) {
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

    fmt.Printf("Buried: [%d] : %s\n", id, body)
}
