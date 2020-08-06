package tree

import (
	"fmt"
	"testing"
)

func TestHuffmanTree(t *testing.T) {
	tree := GetHuffmanTree("beep boop beer!")
	if tree.root.Left.Left.Value != "b" {
		t.Fail()
	}
}

func TestGetSortHuffmanNodesByText(t *testing.T) {
	var texts = []struct {
		Text   string
		Length int
	}{
		{"Hello, world!", 10},
		{"你好，世界", 5},
		{"Bonjour, le monde", 12},
		{"こんにちは、世界", 8},
		{"안녕, 세상", 6},
	}
	for _, text := range texts {
		nodes := getSortHuffmanNodesByText(text.Text)
		if nodes.Len() != text.Length {
			t.Error(fmt.Sprintf("text: %s, length: %d, nodes length: %d", text.Text, text.Length, nodes.Len()))
		}
	}
}

func TestGenerateCodeTable(t *testing.T) {
	tree := GetHuffmanTree("beep boop beer!")
	table := tree.GenerateCodeTable()
	if table["b"] != "00" {
		t.Fail()
	}
}
