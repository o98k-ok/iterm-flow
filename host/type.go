package host

import "container/list"

type Type struct {
	Name string
}

type TypeDependency map[string]*list.List

func NewTypeDependency(deps [][]string) TypeDependency {
	res := make(map[string]*list.List)
	for _, dep := range deps {
		key := dep[len(dep)-1]
		res[key] = list.New()
		for _, d := range dep {
			res[key].PushBack(Type{d})
		}
	}

	return res
}

func (t TypeDependency) GetDependencyLine(typeName string) *list.List {
	return t[typeName]
}