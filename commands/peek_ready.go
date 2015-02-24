package commands

import (
	"fmt"
	"github.com/kr/beanstalk"
	"github.com/spf13/cobra"
	"log"
)

var tubePeekReadyCommand = &cobra.Command{
	Use:   "peek_ready [tube_name]",
	Short: "Retrieve the first available job in the ready queue",
	Long:  `Let's see what jobs are available but lacking reservations.`,
}

func init() {
	tubePeekReadyCommand.Run = tube_peek_ready
}

func tube_peek_ready(cmd *cobra.Command, args []string) {
	client := connect()
	defer client.Close()

	client.Tube = beanstalk.Tube{client, tube_name}

	id, body, err := client.Tube.PeekReady()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ready: [%d] : %s\n", id, body)
}
