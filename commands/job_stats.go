package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "strconv"
)

var jobStatsCommand = &cobra.Command{
    Use:   "job_stats [job_id]",
    Short: "Retrieve statistics about the given job",
    Long:  `Retrieve statistics about the given job!`,
}

func init() {
    jobStatsCommand.Run = job_stats
}

func job_stats(cmd *cobra.Command, args []string) {
    if len(args) < 1 {
        cmd.Help()
        return
    }

    id, err := strconv.ParseUint(args[0], 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    client := connect()
    defer client.Close()

    stats, err := client.StatsJob(id)
    if err != nil {
        log.Fatal(err)
    }

    for k, v := range stats {
        fmt.Printf("%s : %s\n", k, v)
    }
}
