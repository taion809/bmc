package commands

import (
	"fmt"
	"github.com/kr/beanstalk"
	"github.com/spf13/cobra"
	"log"
)

var tubePeekDelayedCommand = &cobra.Command{
	Use:   "peek_delayed [tube_name]",
	Short: "Retrieve the first available job in the delayed queue",
	Long:  `This jobs train was running late.`,
}

func init() {
	tubePeekDelayedCommand.Run = tube_peek_delayed
}

func tube_peek_delayed(cmd *cobra.Command, args []string) {
	client := connect()
	defer client.Close()

	client.Tube = beanstalk.Tube{client, tube_name}

	id, body, err := client.Tube.PeekBuried()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Delayed: [%d] : %s\n", id, body)
}
