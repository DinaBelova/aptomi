package state

import (
	"fmt"
	"github.com/Aptomi/aptomi/pkg/client/rest"
	"github.com/Aptomi/aptomi/pkg/client/rest/http"
	"github.com/Aptomi/aptomi/pkg/config"
	"github.com/spf13/cobra"
)

func newEnforceCommand(cfg *config.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enforce",
		Short: "state enforce",
		Long:  "state enforce long",

		Run: func(cmd *cobra.Command, args []string) {
			rev, err := rest.New(cfg, http.NewClient(cfg)).State().Reset()

			if err != nil {
				panic(fmt.Sprintf("Error while state enforcement: %s", err))
			}

			// todo(slukjanov): replace with -o yaml / json / etc handler
			fmt.Println("Current revision:", rev)
		},
	}

	return cmd
}
