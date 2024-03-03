package docs

import (
	pluginsCommon "github.com/jfrog/jfrog-cli-core/v2/plugins/common"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

const (
	Cmd1 = "cmd-1"
	Cmd2 = "cmd-2"
)

const (
	// Base flags keys
	ServerId    = "server-id"
	url         = "url"
	user        = "user"
	password    = "password"
	accessToken = "access-token"
)

// Security Flag keys mapped to their corresponding components.Flag definition.
var flagsMap = map[string]components.Flag{
	// Common commands flags
	ServerId:    components.NewStringFlag(ServerId, "Server ID configured using the config command."),
	url:         components.NewStringFlag(url, "JFrog Xray URL."),
	user:        components.NewStringFlag(user, "JFrog username."),
	password:    components.NewStringFlag(password, "JFrog password."),
	accessToken: components.NewStringFlag(accessToken, "JFrog access token."),
}
var commandFlags = map[string][]string{
	Cmd1: {
		url, user, password, accessToken, ServerId,
	},
	Cmd2: {
		url, user, password, accessToken, ServerId,
	},
}

func GetCommandFlags(cmdKey string) []components.Flag {
	return pluginsCommon.GetCommandFlags(cmdKey, commandFlags, flagsMap)
}
