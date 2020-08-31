package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get info on object",
	Long:  `get info on object`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provide an object name, e.g. project or model")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
