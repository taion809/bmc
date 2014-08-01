package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "strconv"
)

var touchCommand = &cobra.Command{
    Use:   "touch [job_id]",
    Short: "Reset reservation timer for the given job",
    Long:  `Reset reservation timer for the given job!`,
}

func init() {
    touchCommand.Run = touch
}

func touch(cmd *cobra.Command, args []string) {
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

    err = client.Touch(id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Reset reservation timer for [%d]\n", id)
}
