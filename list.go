package main

import (
	"encoding/json"
	"fmt"
	"github.com/o98k-ok/lazy/app"
	"github.com/o98k-ok/lazy/host"
	"github.com/o98k-ok/lazy/utils"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		node   host.Node
		err    error
		chains [][]string
		nodes = make([]host.Node, 0)
		custom = "custom"
		jumper = "jumper"
	)

	jumperInfo := os.Getenv(jumper)
	if !utils.Empty(jumperInfo) {
		nodes = append(nodes, host.Node{
			Name: jumper,
			Tags: []string{jumper},
			IP:   jumperInfo,
			Type: jumper,
		})
		chains = [][]string{
			[]string{jumper},
			[]string{jumper},
		}
	}

	if len(chains) == 0 {
		chains = make([][]string, 1)
		chains[0] = make([]string, 0)
	}
	chains[len(chains)-1] = append(chains[len(chains)-1], custom)
	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, custom+"_") {
			continue
		}

		es := strings.Split(env, "=")
		if len(es) != 2 {
			continue
		}

		if err = json.Unmarshal([]byte(es[1]), &node); err != nil {
			continue
		}
		node.Name = es[0][7:]
		node.Type = custom
		nodes = append(nodes, node)
	}

	app.SetIconPath("./lazy.jpeg")
	if err = app.InitApp(chains, nodes); err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}

	items := app.FilterByTags(os.Args[1:])
	if len(items.Items) == 0 {
		items.Append(app.NewItem("NotFound", "no such tag info", "", app.IconPath))
	}
	fmt.Println(items.Encode())
}
