package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, "error:", err)
		os.Exit(1)
	}
}

func main() {
	plugin.Start(&TachyonPlugin{})
}

type TachyonPlugin struct{}

func (plugin TachyonPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	spaceName := args[1]

	_, err := cliConnection.CliCommand("create-space", spaceName)
	fatalIf(err)

	_, err = cliConnection.CliCommand("target", "-s", spaceName)
	fatalIf(err)
}

func (TachyonPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "tachyon",
		Commands: []plugin.Command{
			{
				Name:     "take-space",
				HelpText: "Create a space and then switch to it immediately",
			},
		},
	}
}
