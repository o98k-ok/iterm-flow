package core

import (
	"encoding/json"
	"github.com/gookit/validate"
	"github.com/samber/lo"
)

type Node struct {
	Name string   `json:"name" validate:"required"`
	Tags []string `json:"tags" validate:"required"`

	IP     string `json:"ip" validate:"required"`
	Port   string `json:"port" validate:"required"`
	User   string `json:"user" validate:"required"`
	Passwd string `json:"passwd"`
	Depend string `json:"depend"`
}

const (
	DefaultPort = "22"
	DefaultUser = "root"
)

func NewNode(entry lo.Entry[string, string]) *Node {
	node := Node{
		Name: entry.Key,
		Port: DefaultPort,
		User: DefaultUser,
	}

	if err := json.Unmarshal([]byte(entry.Value), &node); err != nil {
		return nil
	}
	extraTags := []string{entry.Key, node.IP, node.Port, node.User}
	node.Tags = append(node.Tags, extraTags...)

	if !validate.Struct(node).Validate() {
		return nil
	}
	return &node
}
