package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Tags   []string `json:"tags"`
	IP     string   `json:"ip"`
	Port   string   `json:"port,omitempty"`
	User   string   `json:"user,omitempty"`
	Passwd string   `json:"passwd,omitempty"`
	Depend string   `json:"depend,omitempty"`
}

////////////////////////////////////////////////
//   Introduce: generate alfred json env
//   Sample:
//      custom_001=> {"tags":["abc", "dev"], "ip":"10.10.10.10", "user":"ok"}
//   Action:
//		1. ssh jumper@xxx.com
//		2. ssh ok@10.10.10.10
////////////////////////////////////////////////
func main() {
	nodes := make(map[string]Node)
	reader := bufio.NewReader(os.Stdin)

	var input string
	for {
		input = Read(reader, "Add a custom node[Y/N]? ")
		if input == "N" {
			break
		}

		name := Read(reader, "[Required]Input node name: ")
		nodes[name] = Node{
			Tags:   strings.Fields(Read(reader, "[Required]Input tags: ")),
			IP:     Read(reader, "[Required]Input login ip/host: "),
			Port:   Read(reader, "[Optional]Input port: "),
			User:   Read(reader, "[Optional]Input username: "),
			Passwd: Read(reader, "[Optional]Input password: "),
			Depend: Read(reader, "[Optional]Input depend name: "),
		}
	}

	fmt.Println("////////////////////////////////////////////////")
	fmt.Println("Please copy these env into alfred:")
	for key, value := range nodes {
		d, err := json.Marshal(value)
		if err != nil {
			continue
		}
		fmt.Printf("Name:%s Value:%s\n", key, string(d))
	}
	fmt.Println("////////////////////////////////////////////////")
}

func Read(reader *bufio.Reader, notice string) string {
	fmt.Print(notice)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(input)
}
