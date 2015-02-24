package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var connStatsCommand = &cobra.Command{
	Use:   "stats",
	Short: "Retrieve statistics about the given connection",
	Long:  `Retrieve statistics about the given connection!`,
}

func init() {
	connStatsCommand.Run = conn_stats
}

func conn_stats(cmd *cobra.Command, args []string) {
	client := connect()
	defer client.Close()

	stats, err := client.Stats()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range stats {
		fmt.Printf("%s : %s\n", k, v)
	}
}
