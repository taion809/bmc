package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "strconv"
)

var buryCommand = &cobra.Command{
    Use:   "bury [job_id] [priority]",
    Short: "Bury [job_id] with [priority]",
    Long: `Bury [job_id] with [priority].  
This job will remain in the buried state until kicked.  Job must currently be reserved or beanstalkd will return NOT_FOUND.`,
}

func init() {
    buryCommand.Run = bury_job
}

func bury_job(cmd *cobra.Command, args []string) {
    if len(args) < 2 {
        cmd.Help()
        return
    }

    id, err := strconv.ParseUint(args[0], 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    pri64, err := strconv.ParseUint(args[1], 10, 32)
    if err != nil {
        log.Fatal(err)
    }

    pri := uint32(pri64)

    client := connect()
    defer client.Close()

    err = client.Bury(id, pri)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Buried job: [%d] with priority [%d]\n", id, pri)
}
