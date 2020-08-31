package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an object",
	Long:  `create an object`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provide an object name, e.g. project or model")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
