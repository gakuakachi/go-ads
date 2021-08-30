package main

type UnionFind struct {
	arr []int
}

func (uf *UnionFind) Len() int {
	return len(uf.arr)
}

func (uf *UnionFind) Size(nd int) int {
	return -uf.arr[uf.Root(nd)]
}

func (uf *UnionFind) Root(nd int) int {
	if uf.arr[nd] < 0 {
		return nd
	}

	return uf.Root(uf.arr[nd])
}

func (uf *UnionFind) Unite(a, b int) bool {
	ar := uf.Root(a)
	br := uf.Root(b)

	if ar == br {
		return false
	}

	asize := uf.Size(a)
	bsize := uf.Size(b)

	if bsize > asize {
		ar, br = br, ar
		bsize = asize
	}

	uf.arr[ar] -= bsize
	uf.arr[br] = ar

	return true
}
