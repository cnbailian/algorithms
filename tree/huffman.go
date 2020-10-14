package tree

import (
	"sort"
)

type HuffmanTree struct {
	root *HuffmanNode
}

type HuffmanNode struct {
	Value string
	Count int
	Left  *HuffmanNode
	Right *HuffmanNode
}

func GetHuffmanTree(text string) HuffmanTree {
	nodes := getSortHuffmanNodesByText(text)
	return HuffmanTree{
		root: HuffmanNodesSort(nodes),
	}
}

func (tree HuffmanTree) GenerateCodeTable() HuffmanCodeTable {
	table := HuffmanCodeTable{}
	TraverseHuffmanTree(tree.root, "", table)
	return table
}

type SortHuffmanNodes []*HuffmanNode

func (s SortHuffmanNodes) Len() int {
	return len(s)
}

func (s SortHuffmanNodes) Less(i, j int) bool {
	return s[i].Count < s[j].Count
}

func (s SortHuffmanNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func HuffmanNodesSort(nodes SortHuffmanNodes) *HuffmanNode {
	sort.Sort(nodes)
	if nodes.Len() == 0 {
		return nil
	}
	if nodes.Len() > 1 {
		newNodes := append(nodes[2:], &HuffmanNode{
			Count: nodes[0].Count + nodes[1].Count,
			Left:  nodes[0],
			Right: nodes[1],
		})
		return HuffmanNodesSort(newNodes)
	}
	return nodes[0]
}

func getSortHuffmanNodesByText(text string) SortHuffmanNodes {
	chars := map[string]int{}
	for _, s := range text {
		if _, ok := chars[string(s)]; ok {
			chars[string(s)] += 1
		} else {
			chars[string(s)] = 1
		}
	}
	nodes := SortHuffmanNodes{}
	for value, count := range chars {
		nodes = append(nodes, &HuffmanNode{
			Value: value,
			Count: count,
		})
	}
	return nodes
}

type HuffmanCodeTable map[string]string

func TraverseHuffmanTree(node *HuffmanNode, code string, table HuffmanCodeTable) {
	if node == nil {
		return
	}
	if node.Value != "" {
		table[node.Value] = code
	}
	TraverseHuffmanTree(node.Left, code+"0", table)
	TraverseHuffmanTree(node.Right, code+"1", table)
}
