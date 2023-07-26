package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/samber/lo"
)

func EveryContains(collection []string, subset []string) bool {
	for _, elem := range subset {
		if !lo.ContainsBy(collection, func(t string) bool {
			return strings.Contains(t, elem)
		}) {
			return false
		}
	}

	return true
}

func Select(variables map[string]string, key []string) []*Node {
	var nodes []*Node

	for _, entry := range lo.Entries[string, string](variables) {
		node := NewNode(entry)
		if node != nil && EveryContains(node.Tags, key) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func Trace(variables map[string]string, key string) []*Node {
	var nodes []*Node

	depends := lo.MapValues[string, string, *Node](variables, func(v string, k string) *Node {
		return NewNode(lo.Entry[string, string]{Key: k, Value: v})
	})

	for len(key) != 0 {
		node, ok := depends[key]
		if !ok {
			break
		}
		nodes = append([]*Node{node}, nodes...)
		key = node.Depend
	}
	return nodes
}

func GetField(variables map[string]string, field string, key string) string {
	val, ok := variables[key]
	if !ok {
		return "404"
	}

	var node map[string]interface{}
	err := json.Unmarshal([]byte(val), &node)
	if err != nil {
		return "ERROR"
	}
	return node[field].(string)
}

func Display(nodes []*Node) *alfred.Items {
	res := alfred.NewItems()
	lo.ForEach(nodes, func(t *Node, _ int) {
		item := alfred.NewItem(t.Name, fmt.Sprintf("%s-[%s]", t.IP, strings.Join(t.Tags, "|")), t.Name)
		res.Append(item)
	})
	return res
}

func Routine(nodes []*Node) string {
	if len(nodes) <= 0 {
		return ""
	}

	var arr []string
	for _, n := range nodes {
		cmd := fmt.Sprintf("ssh %s %s@%s -p%s", n.PrefixExtra, n.User, n.IP, n.Port)
		if len(n.Passwd) != 0 {
			cmd += fmt.Sprintf(";%s", n.Passwd)
		}

		arr = append(arr, cmd)
	}

	return fmt.Sprintf("%s#%s", nodes[len(nodes)-1].Name, strings.Join(arr, ";"))
}
