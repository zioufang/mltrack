package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "create a project",
	Long:  `create a project`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		project := model.Project{
			Name:        name,
			Description: description,
		}
		reqBody, err := json.Marshal(project)
		if err != nil {
			log.Fatalln(err)
		}
		resp, err := http.Post(URI+projectEndPoint, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("project created:")
		fmt.Println(string(respBody))
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)
	createProjectCmd.Flags().StringP("name", "n", "", "name for the project")
	createProjectCmd.MarkFlagRequired("name")
	createProjectCmd.Flags().StringP("description", "d", "", "description for the project")
}
