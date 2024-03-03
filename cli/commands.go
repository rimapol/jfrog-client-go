package cli

import (
	flags "github.com/jfrog/jfrog-cli-artifactory/cli/docs"
	cmd1docs "github.com/jfrog/jfrog-cli-artifactory/cli/docs/cmd1"
	cmd2docs "github.com/jfrog/jfrog-cli-artifactory/cli/docs/cmd2"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

const category = "My Category"

func getCommands() []components.Command {
	return []components.Command{
		{
			Name:        "cmd-1",
			Aliases:     []string{"c1"},
			Flags:       flags.GetCommandFlags(flags.Cmd1),
			Description: cmd1docs.GetDescription(),
			Arguments:   cmd1docs.GetArguments(),
			Category:    category,
			Action:      cmd1,
		},
		{
			Name:        "cmd-2",
			Aliases:     []string{"c2"},
			Flags:       flags.GetCommandFlags(flags.Cmd2),
			Description: cmd2docs.GetDescription(),
			Arguments:   cmd2docs.GetArguments(),
			Category:    category,
			Action:      cmd2,
		},
	}
}

func cmd1(c *components.Context) error {
	return nil
}

func cmd2(c *components.Context) error {
	return nil
}
