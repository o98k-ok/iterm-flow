package host

type Node struct {
	Name string
	IP string
	Port string
	User string
	Passwd string
	Tags []string
	Type string
}


type NodeGraph struct {
	// all nodes type relations
	Types TypeDependency
	// all nodes list: key=ip
	Nodes map[string]*Node
	// key: node one tag
	Tags map[string][]*Node
}

func NewNodeGraph() *NodeGraph {
	return nil
}

func ListNodeByTags(tags []string)  []*Node {
	return nil
}

// ListNodeByKey
//  @param key
// 		substr of Node Name or Node IP
//  @return []*Node
// 		node list
func ListNodeByKey(key string)  []*Node {
	return nil
}

func GenNodeRelation(node *Node) []*Node {
	return nil
}