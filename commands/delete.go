package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var deleteCommand = &cobra.Command{
	Use:   "delete [job_id]",
	Short: "Delete [job_id]",
	Long: `Delete [job_id].  
Remove [job_id] from beanstalkd.`,
}

func init() {
	deleteCommand.Run = delete_job
}

func delete_job(cmd *cobra.Command, args []string) {
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

	err = client.Delete(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted job: [%d]\n", id)
}
