package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// URI is the global uri for this mltrack instance
// Using the MLTRACK_URI env var, if not provided default to localhost:8000
var URI = getURI()

// TODO move these to api package and import from there, also for subroutes e.g. /projects/all
// projectEndPoint is the root api endpoint for projects
const projectEndPoint = "/projects"

// projectEndPoint is the root api endpoint for models
const modelEndPoint = "/models"

// projectEndPoint is the root api endpoint for model runs
const modelRunEndPoint = "/runs"

var rootCmd = &cobra.Command{
	Use:   "mltrack",
	Short: "mltrack tracks analytics models",
	Long:  `A lightweight alternative to MLflow built with the focus on model tracking`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("just chilling on " + URI)
	},
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getURI() string {
	var uri string
	var ok bool
	if uri, ok = os.LookupEnv("MLTRACK_URI"); !ok {
		// TODO parameterize the default uri toghether with the api
		uri = "http://localhost:8000"
	}
	return uri
}
