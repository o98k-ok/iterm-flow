package core

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	entry := lo.Entry[string, string]{Key: "server01"}

	t.Run("test default value", func(t *testing.T) {
		entry.Value = "{\"ip\":\"127.0.0.1\"}"
		expect := Node{
			Name: "server01",
			Tags: []string{"server01", "127.0.0.1", "22", "root"},
			IP:   "127.0.0.1",
			Port: "22",
			User: "root",
		}
		node := NewNode(entry)
		assert.Equal(t, expect, *node)
	})

	t.Run("test with tags", func(t *testing.T) {
		entry.Value = "{\"tags\":[\"dev\"],\"ip\":\"127.0.0.1\",\"user\":\"root\",\"passwd\":\"oh-amazing\"}"
		expect := Node{
			Name:   "server01",
			Tags:   []string{"dev", "server01", "127.0.0.1", "22", "root"},
			IP:     "127.0.0.1",
			Port:   "22",
			User:   "root",
			Passwd: "oh-amazing",
		}

		node := NewNode(entry)
		assert.Equal(t, expect, *node)
	})

	t.Run("test without ip", func(t *testing.T) {
		entry.Value = "{\"tags\":[\"dev\"]}"
		node := NewNode(entry)
		assert.Nil(t, node)
	})

	t.Run("test json format error", func(t *testing.T) {
		entry.Value = ""
		node := NewNode(entry)
		assert.Nil(t, node)
	})
}
