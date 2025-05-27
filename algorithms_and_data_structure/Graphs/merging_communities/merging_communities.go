package main

import "fmt"

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	repX := uf.Find(x)
	repY := uf.Find(y)
	if repX != repY {
		if uf.size[repX] > uf.size[repY] {
			uf.parent[repY] = repX
			uf.size[repX] += uf.size[repY]
		} else {
			uf.parent[repX] = repY
			uf.size[repY] += uf.size[repX]
		}
	}
}

func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

type MergingCommunities struct {
	uf *UnionFind
}

func NewMergingCommunities(n int) *MergingCommunities {
	return &MergingCommunities{uf: NewUnionFind(n)}
}

func (mc *MergingCommunities) Connect(x, y int) {
	mc.uf.Union(x, y)
}

func (mc *MergingCommunities) GetCommunitySize(x int) int {
	return mc.uf.GetSize(x)
}

// Example usage
func main() {
	mc := NewMergingCommunities(5)
	mc.Connect(0, 1)
	mc.Connect(1, 2)
	fmt.Println(mc.GetCommunitySize(0)) // Output: 3
	fmt.Println(mc.GetCommunitySize(3)) // Output: 1
}
