package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var getProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "get project(s)",
	Long:  `get project(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		var resp *http.Response
		var err error

		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		if id != "" {

		} else if name != "" {

		} else {
			resp, err = http.Get(URI + projectEndPoint + "/all")
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(respBody))
	},
}

func init() {
	getCmd.AddCommand(getProjectCmd)
	getProjectCmd.Flags().StringP("id", "i", "", "id for the project")
	getProjectCmd.Flags().StringP("name", "n", "", "name for the project")
}
