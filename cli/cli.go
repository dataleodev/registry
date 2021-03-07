package cli

import "github.com/spf13/cobra"

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

func MakeAllCommands() Commands {
	return Commands{
		Login:    nil,
		Register: nil,
		Users:    nil,
		Nodes:    nil,
		Regions:  nil,
	}
}
