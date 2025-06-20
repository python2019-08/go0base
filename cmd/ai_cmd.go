package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	aiTrans "go0base/ai/translator"
)

type tAiCommand struct {
	cmd *cobra.Command

	appName string
}

func newAiCommand() *tAiCommand {
	ret := &tAiCommand{
		appName: "go0base",
	}

	ret.cmd = &cobra.Command{
		Use:     "ai",
		Short:   "start a AI app",
		Long:    "Use when you need to create a new AI app",
		Example: "go0base ai -n admin",
		Run: func(cmd *cobra.Command, args []string) {
			ret.run()
		},
	}

	return ret
}

func (g *tAiCommand) init() {
	g.cmd.PersistentFlags().StringVarP(&(g.appName), "name", "n", "", "Start server with provided configuration file")
}

func (g *tAiCommand) run() {
	fmt.Println(`tAiCommand.run()...start`)
	defer fmt.Println(`tAiCommand.run()...end`)

	aiTrans.TransLator_main()
}
