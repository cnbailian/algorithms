package graph

import (
	"testing"
)

func getAsciiGraph() *asciiGraph {
	ag := &asciiGraph{Title: "test"}
	ag.init()
	return ag
}

func TestAsciiGraph_AddCollectionRow(t *testing.T) {
	ag := getAsciiGraph()
	ag.AddCollectionRow("collection", []string{"a", "b", "c"})
	expected := `|  test
|  
|  collection: collection
|  [a, b, c]
`
	if ag.result() != expected {
		t.Fail()
	}
}

func TestAsciiGraph_AddTableRows(t *testing.T) {
	ag := getAsciiGraph()
	ag.AddTableRows("test", [][]string{
		{"0", "1", "1"},
		{"1", "0", "1"},
		{"1", "1", "0"},
	})
	expected := `|  test
|  
|  table: test
|  0 1 1 
|  1 0 1 
|  1 1 0 
`
	if ag.result() != expected {
		t.Fail()
	}
}
