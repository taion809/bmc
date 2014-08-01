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
    client := connect()
    defer client.Close()

    client.Tube = beanstalk.Tube{client, tube_name}
    id, err := client.Put([]byte("hello"), 1, 0, 3800*time.Second)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("[x] Sent: ", id)
}
