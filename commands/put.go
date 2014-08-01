package commands

import (
    "fmt"
    "github.com/kr/beanstalk"
    "github.com/spf13/cobra"
    "io/ioutil"
    "log"
    "time"
)

var tubePutCommand = &cobra.Command{
    Use:   "put",
    Short: "put",
    Long:  "put",
}

func init() {
    tubePutCommand.Flags().Int64Var(&ttr, "ttr", 120, "Sets the TTR for the job")
    tubePutCommand.Flags().Int64Var(&delay, "delay", 0, "Sets the delay timer for the job")
    tubePutCommand.Flags().Uint32Var(&pri, "pri", 100, "Sets job priority")
    tubePutCommand.Flags().StringVarP(&filename, "filename", "f", "", "Read in a file as the job body")
    tubePutCommand.Run = tube_put
}

func tube_put(cmd *cobra.Command, args []string) {
    if len(args) < 1 && len(filename) == 0 {
        cmd.Help()
        return
    }

    var body []byte

    if len(args) > 0 {
        body = []byte(args[0])
    } else {
        fileBody, err := ioutil.ReadFile(filename)
        if err != nil {
            log.Fatal(err)
        }

        body = fileBody
    }

    client := connect()
    defer client.Close()

    client.Tube = beanstalk.Tube{client, tube_name}

    delayDuration := time.Duration(delay) * time.Second
    ttrDuration := time.Duration(ttr) * time.Second
    id, err := client.Tube.Put(body, pri, delayDuration, ttrDuration)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("[x] Sent: ", id)
}
