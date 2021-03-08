package cli

import (
	"context"
	"fmt"
	"github.com/dataleodev/registry/api"
	"github.com/spf13/cobra"
	"os"
)


type Command int

const (
	Login Command = iota
	Register
	Users
	Nodes
	Regions
)

//RunFunc wraps the run func in cobra.Command
type RunFunc func(cmd *cobra.Command, args []string)

type CommandRunner interface {
	Run(command Command) RunFunc
}

type runner struct {
	endpoints api.Endpoints
}

var _ CommandRunner = (*runner)(nil)

func NewCommandRunner(endpoints api.Endpoints) CommandRunner {
	return runner{endpoints: endpoints}
}

func (r runner) Run(command Command) RunFunc {
	switch command {
	case Login:
		return r.runLoginCommand()

	default:
		return func(cmd *cobra.Command, args []string) {
			fmt.Printf("this should not happen\n")
		}
	}

}



type Commands struct {
	Login    *cobra.Command
	Register *cobra.Command
	Users    *cobra.Command
	Nodes    *cobra.Command
	Regions  *cobra.Command
	Version  *cobra.Command
	DB       *cobra.Command
	About    *cobra.Command
	Events   *cobra.Command
	Config   *cobra.Command
}

func MakeAllCommands(r CommandRunner) Commands {
	return Commands{
		Login:    makeLoginCommand(r),
		Register: nil,
		Users:    nil,
		Nodes:    nil,
		Regions:  nil,
	}
}

func makeLoginCommand(r CommandRunner) *cobra.Command {
	// loginCmd represents the login command
	var loginCmd = &cobra.Command{
		Use:     "login",
		Short:   "generate auth token",
		Example: "regctl login -u <username> -p <password>",
		Long:    `generates a jwt token string after the user has supplied correct username and password`,
		Run:     r.Run(Login),
	}

	loginCmd.PersistentFlags().StringP("username","u","","username")
	loginCmd.PersistentFlags().StringP("password","p","","password")

	return loginCmd
}

func (r *runner) runLoginCommand()RunFunc {
	return func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		password, err := cmd.Flags().GetString("password")

		if err != nil || username == "" || password == "" {
			print("no username or password")
			os.Exit(1)
		}

		token, err := r.endpoints.Login(context.Background(), username, password)

		if err != nil{
			fmt.Printf("could not generate token %s\n",err.Error())
			os.Exit(1)
		}

		fmt.Printf("token: %s\n",token)
		return

	}
}
