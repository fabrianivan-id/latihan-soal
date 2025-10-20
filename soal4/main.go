package main

import (
	"sort"
)

type edge struct{ u, v, w int }

type dsu struct{ p, r []int }

func newDSU(n int) *dsu {
	p := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	return &dsu{p, r}
}
func (d *dsu) find(x int) int {
	if d.p[x] != x {
		d.p[x] = d.find(d.p[x])
	}
	return d.p[x]
}
func (d *dsu) unite(a, b int) bool {
	pa, pb := d.find(a), d.find(b)
	if pa == pb {
		return false
	}
	if d.r[pa] < d.r[pb] {
		pa, pb = pb, pa
	}
	d.p[pb] = pa
	if d.r[pa] == d.r[pb] {
		d.r[pa]++
	}
	return true
}

func GetMinimumCostMST(graph_nodes int, graph_from, graph_to, graph_weight []int, source, destination int) int {
	m := len(graph_from)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		edges[i] = edge{graph_from[i], graph_to[i], graph_weight[i]}
	}
	// Build MST by Kruskal
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })
	d := newDSU(graph_nodes)
	adj := make([][]edge, graph_nodes+1)
	for _, e := range edges {
		if d.unite(e.u, e.v) {
			adj[e.u] = append(adj[e.u], edge{e.u, e.v, e.w})
			adj[e.v] = append(adj[e.v], edge{e.v, e.u, e.w})
		}
	}

	// If source and destination not connected in MST, return -1
	if d.find(source) != d.find(destination) {
		return -1
	}

	// Unique path in tree: DFS sum of weights
	seen := make([]bool, graph_nodes+1)
	var sum int
	var found bool
	var dfs func(int, int) error
	dfs = func(u, target int) error {
		if seen[u] {
			return nil
		}
		seen[u] = true
		if u == target {
			found = true
			return nil
		}
		for _, nb := range adj[u] {
			if !seen[nb.v] {
				sum += nb.w
				if err := dfs(nb.v, target); err != nil {
					return err
				}
				if found {
					return nil
				}
				sum -= nb.w // backtrack
			}
		}
		return nil
	}
	_ = dfs(source, destination)
	if !found {
		return -1
	}
	return sum
}
