package main

import (
	"encoding/json"
	"fmt"
	"github.com/o98k-ok/lazy/app"
	"github.com/o98k-ok/lazy/host"
	"github.com/o98k-ok/lazy/plist"
	"io"
	"os"
)

const InfoPlist = "info.plist"
const IconPath = "./lazy.jpeg"

func main() {
	var (
		node   host.Node
		flow   *plist.Info
		err    error
		nodes = make([]host.Node, 0)
	)

	if flow, err = plist.NewPlist(InfoPlist); err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}

	vars := flow.GetAttrByNames([][]string{[]string{"variables"}})
	if len(vars) != 1 {
		io.WriteString(os.Stderr, "read variables error"+"\n")
		return
	}

	envs, ok := vars[0].(map[string]interface{})
	if !ok {
		io.WriteString(os.Stderr, "parse variables error"+"\n")
		return
	}

	for name, value := range envs {
		if err = json.Unmarshal([]byte(value.(string)), &node); err != nil {
			continue
		}
		node.Name = name
		nodes = append(nodes, node)
	}

	app.SetIconPath(IconPath)
	if err = app.InitApp(nodes); err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}

	items := app.FilterByTags(os.Args[1:])
	if len(items.Items) == 0 {
		items.Append(app.NewItem("NotFound", "no such tag info", "", app.IconPath))
	}
	fmt.Println(items.Encode())
}
