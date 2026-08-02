package daemon

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var ConfigFile string
var config ConfigStruct

func Start(cmd *cobra.Command, args []string) {
	f, err := os.Open(ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err)
	}
}
