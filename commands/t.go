package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "log"
    "time"
)

var debugCmd = &cobra.Command{
    Use: "debug",
}

func init() {
    debugCmd.Run = add
}

func add(cmd *cobra.Command, args []string) {
    //conf := InitializeConfig()

    client, err := beanstalk.Dial(protocol, hostname+":"+port)
    if err != nil {
        log.Fatal(err)
    }

    defer client.Close()

    id, err := client.Put([]byte("hello"), 1, 0, 120*time.Second)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("[x] Sent: ", id)
}
