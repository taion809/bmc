package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var tubeStatsCommand = &cobra.Command{
    Use:   "stats [tube_name]",
    Short: "Retrieve statistics about the given tube",
    Long:  `Refreshing statistics about your favorite tube!`,
}

func init() {
    tubeStatsCommand.Run = tube_stats
}

func tube_stats(cmd *cobra.Command, args []string) {
    client, err := beanstalk.Dial(protocol, hostname+":"+port)
    if err != nil {
        log.Fatal(err)
    }

    defer client.Close()

    client.Tube = beanstalk.Tube{client, tube_name}

    stats, err := client.Tube.Stats()
    if err != nil {
        log.Fatal(err)
    }

    for k, v := range stats {
        fmt.Printf("%s : %s\n", k, v)
    }
}
