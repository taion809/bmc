package commands

import (
	"fmt"
	"github.com/kr/beanstalk"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var tubePauseCommand = &cobra.Command{
	Use:   "pause",
	Short: "pause",
	Long:  "pause",
}

func init() {
	tubePauseCommand.Run = tube_pause
}

func tube_pause(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Help()
		return
	}

	duration, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	client := connect()
	defer client.Close()

	client.Tube = beanstalk.Tube{client, tube_name}

	durDuration := time.Duration(duration) * time.Second
	err = client.Tube.Pause(durDuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Paused %s tube for %d seconds.\n", tube_name, duration)
}
