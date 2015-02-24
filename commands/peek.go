package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var peekCommand = &cobra.Command{
	Use:   "peek [job_id]",
	Short: "peek [job_id]",
	Long:  `Obtain [job_id] contents.`,
}

func init() {
	peekCommand.Run = peek_job
}

func peek_job(cmd *cobra.Command, args []string) {
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

	body, err := client.Peek(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}
