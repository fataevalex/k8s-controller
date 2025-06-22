package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var goBasicCmd = &cobra.Command{
	Use:   "go-basic",
	Short: "Run golang basic code",
	Run: func(cmd *cobra.Command, args []string) {
		//Go basic code to run functions
		log.Debug().Msg("go-basic command starting")
		k8s := Kubernetes{
			Name:    "k8s-demo-cluster",
			Version: "1.31",
			Users:   []string{"alex", "den"},
			NodeNumber: func() int {
				return 10
			},
		}
		log.Info().Str("cluster", k8s.Name).Str("version", k8s.Version).Msg("Initialized Kubernetes cluster")

		//print users
		log.Info().Msg("Getting current users")
		k8s.GetUsers()

		//add new user to struct
		log.Info().Str("user", "anonymous").Msg("Adding new user")
		k8s.AddNewUser("anonymous")

		//print users one more time
		log.Info().Msg("Getting updated users list")
		k8s.GetUsers()

		log.Debug().Msg("go-basic command completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(goBasicCmd)

}

// My go basic fucntions here
type Kubernetes struct {
	Name       string     `json:"name"`
	Version    string     `json:"version"`
	Users      []string   `json:"users,omitempty"`
	NodeNumber func() int `json:"-"`
}

func (k8s Kubernetes) GetUsers() {
	for _, user := range k8s.Users {
		log.Debug().Msg("Starting GetUsers command")
		fmt.Println(user)
	}
}

func (k8s *Kubernetes) AddNewUser(user string) {
	log.Debug().Msg("Starting AddNewUser command")
	k8s.Users = append(k8s.Users, user)
}
