package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
)

var listTubesCmd = &cobra.Command{
    Use:   "list_tubes",
    Short: "List the active tubes",
    Long:  `List the active tubes`,
}

func init() {
    listTubesCmd.Run = list_tubes
}

func list_tubes(cmd *cobra.Command, args []string) {
    client, err := beanstalk.Dial(protocol, hostname+":"+port)
    if err != nil {
        log.Fatal(err)
    }

    defer client.Close()

    tubes, err := client.ListTubes()
    if err != nil {
        log.Fatal(err)
    }

    for _, tube := range tubes {
        fmt.Println(tube)
    }
}
