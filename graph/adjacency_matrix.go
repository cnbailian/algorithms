package graph

// 邻接矩阵图
type AdjacencyMatrixUnweightedGraph struct {
	Vertexes []string
	Edges    [][]string
}

func (am *AdjacencyMatrixUnweightedGraph) Graph() {
	ag := NewAsciiGraph("Adjacency Matrix Graph")
	ag.AddCollectionRow("vertex", am.Vertexes)
	ag.AddTableRows("edge", am.Edges)
	ag.Print()
}

func (am *AdjacencyMatrixUnweightedGraph) GetVertexDegree(i int) int {
	d := 0
	for _, b := range am.Edges[i] {
		if b == "1" {
			d += 1
		}
	}
	return d
}

func (am *AdjacencyMatrixUnweightedGraph) GetEdge(vi, vj int) bool {
	return am.Edges[vi][vj] == "1" && am.Edges[vj][vi] == "1"
}

func (am *AdjacencyMatrixUnweightedGraph) GetPath(vi, vj int) []string {
	return nil
}

func (am *AdjacencyMatrixUnweightedGraph) IsSimplePath(vi, vj int) []string {
	return nil
}

func (am *AdjacencyMatrixUnweightedGraph) IsCycle() bool {
	return false
}

func (am *AdjacencyMatrixUnweightedGraph) IsSimpleCycle() bool {
	return false
}

func (am *AdjacencyMatrixUnweightedGraph) IsConnectedGraph() bool {
	return false
}

// 加权邻接矩阵图 边值为权重
type AdjacencyMatrixWeightGraph struct {
	Vertexes []string
	Edges    [][]float64
}
