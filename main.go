package main

import (
	"fmt"
	"github.com/o98k-ok/iterm/core"
	"github.com/o98k-ok/lazy/v2/alfred"
	"os"
)

func main() {
	variables, err := alfred.FlowVariables()
	if err != nil {
		alfred.ErrItems("Read variables error", err)
		return
	}

	xli := alfred.NewApp("manage iterm profiles")
	xli.Bind("select", func(patterns []string) {
		if len(patterns) == 0 {
			patterns = []string{""}
		}
		nodes := core.Select(variables, patterns)
		fmt.Println(core.Display(nodes).Encode())
	})

	xli.Bind("trace", func(patterns []string) {
		if len(patterns) < 1 {
			return
		}

		nodes := core.Trace(variables, patterns[0])
		fmt.Println(core.Routine(nodes))
	})

	err = xli.Run(os.Args)
	if err != nil {
		alfred.ErrItems("run failed", err)
	}
}
