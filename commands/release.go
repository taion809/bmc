package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var releaseCommand = &cobra.Command{
	Use:   "release",
	Short: "release",
	Long:  "release",
}

func init() {
	releaseCommand.Flags().Int64Var(&delay, "delay", 0, "Sets the TTR for the job")
	releaseCommand.Flags().Uint32Var(&pri, "pri", 100, "Sets job priority")
	releaseCommand.Run = release_job
}

func release_job(cmd *cobra.Command, args []string) {
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

	delayDuration := time.Duration(delay) * time.Second
	err = client.Release(id, pri, delayDuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[x] Released: ", id)
}
