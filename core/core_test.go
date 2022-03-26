package core

import (
	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSelect(t *testing.T) {
	variables := map[string]string{
		"server01": "{\"tags\":[\"dev\"],\"ip\":\"127.0.0.1\",\"user\":\"root\",\"passwd\":\"oh-amazing\"}",
		"server02": "{\"tags\":[\"dev\"],\"ip\":\"127.0.0.1\",\"user\":\"root\",\"passwd\":\"oh-amazing\"}",
	}
	t.Run("test select name", func(t *testing.T) {
		name := []string{"server01"}
		nodes := Select(variables, name)
		assert.Equal(t, 1, len(nodes))
		assert.Equal(t, name, []string{nodes[0].Name})
	})

	t.Run("test select tags", func(t *testing.T) {
		name := []string{"de"}
		nodes := Select(variables, name)
		assert.Equal(t, 2, len(nodes))

		expect := lo.Keys[string, string](variables)
		got := []string{nodes[0].Name, nodes[1].Name}
		sort.Strings(got)
		sort.Strings(expect)

		assert.Equal(t, expect, got)
	})

	t.Run("test select union", func(t *testing.T) {
		name := []string{"d", "server01"}
		nodes := Select(variables, name)
		assert.Equal(t, 1, len(nodes))
		assert.Equal(t, []string{"server01"}, []string{nodes[0].Name})
	})

	t.Run("test select union", func(t *testing.T) {
		name := []string{""}
		nodes := Select(variables, name)
		assert.Equal(t, 2, len(nodes))

		expect := lo.Keys[string, string](variables)
		got := []string{nodes[0].Name, nodes[1].Name}
		sort.Strings(got)
		sort.Strings(expect)

		assert.Equal(t, expect, got)
	})
}

func TestTrace(t *testing.T) {
	variables := map[string]string{
		"server01": "{\"ip\":\"127.0.0.1\"}",
		"server02": "{\"ip\":\"127.0.0.1\", \"depend\":\"server01\"}",
	}

	t.Run("test single trace", func(t *testing.T) {
		routine := Trace(variables, "server01")
		assert.Equal(t, 1, len(routine))
		assert.Equal(t, "server01", routine[0].Name)
	})

	t.Run("test routine trace", func(t *testing.T) {
		routine := Trace(variables, "server02")
		assert.Equal(t, 2, len(routine))
		expect := []string{"server01", "server02"}

		assert.Equal(t, expect, lo.Map[*Node, string](routine, func(n *Node, _ int) string {
			return n.Name
		}))
	})

	t.Run("test empty trace", func(t *testing.T) {
		routine := Trace(variables, "server03")
		assert.Equal(t, 0, len(routine))
	})
}

func TestMerge(t *testing.T) {
	variables := map[string]string{
		"server01": "{\"ip\":\"127.0.0.1\"}",
		"server02": "{\"ip\":\"127.0.0.1\", \"depend\":\"server01\"}",
	}
	t.Run("test display", func(t *testing.T) {
		nodes := Select(variables, []string{"server"})
		items := Display(nodes)

		expect := []string{"server01", "server02"}
		gotTitle := lo.Map(items.Items, func(f *alfred.Item, _ int) string {
			return f.Title
		})
		sort.Strings(gotTitle)
		gotArg := lo.Map(items.Items, func(f *alfred.Item, _ int) string {
			return f.Arg
		})
		sort.Strings(gotArg)

		assert.Equal(t, 2, items.Len())
		assert.Equal(t, expect, gotArg)
		assert.Equal(t, expect, gotTitle)
	})

	t.Run("test routine", func(t *testing.T) {
		nodes := Trace(variables, "server02")

		expect := "server02#ssh root@127.0.0.1 -p22;ssh root@127.0.0.1 -p22"
		got := Routine(nodes)
		assert.Equal(t, expect, got)
	})

}
