package graph

import "fmt"

type asciiGraph struct {
	Title string
	Table [][]string
}

func (a *asciiGraph) init() {
	if len(a.Table) > 0 {
		a.Table[0] = []string{"error: you must first call the init func"}
		return
	}
	a.Table = append(a.Table, []string{a.Title})
}

func (a *asciiGraph) AddCollectionRow(name string, col []string) {
	a.Table = append(a.Table, []string{}, []string{"collection: " + name})
	row := []string{"["}
	for i, s := range col {
		if i != len(col)-1 {
			s += ", "
		}
		row = append(row, s)
	}
	row = append(row, "]")
	a.Table = append(a.Table, row)
}

func (a *asciiGraph) AddTableRows(name string, table [][]string) {
	a.Table = append(a.Table, []string{}, []string{"table: " + name})
	for _, row := range table {
		var r []string
		for _, s := range row {
			r = append(r, s+" ")
		}
		a.Table = append(a.Table, r)
	}
}

func (a *asciiGraph) result() (res string) {
	for _, horizontal := range a.Table {
		res += "|  "
		for _, v := range horizontal {
			res += v
		}
		res += "\n"
	}
	return
}

func (a *asciiGraph) Print() {
	fmt.Print(a.result())
}
