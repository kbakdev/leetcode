package _go

type BookMyShow struct {
	rows, seats   int
	segmentedTree []segNode2286
}

type segNode2286 struct {
	max int
	sum int
}

func Constructor(rows int, seats int) BookMyShow {
	size := 1
	for size < rows*2 {
		size = size * 2
	}

	obj := BookMyShow{rows, seats, make([]segNode2286, size)}

	obj.build(0, 0, rows-1)

	return obj
}

func (receiver *BookMyShow) build(idx, ldx, rdx int) {
	if ldx == rdx {
		receiver.segmentedTree[idx] = segNode2286{receiver.seats, receiver.seats}
		return
	}

	mdx := (ldx + rdx) / 2
	receiver.segmentedTree[idx] = segNode2286{receiver.seats, (rdx - ldx + 1) * receiver.seats}

	receiver.build(2*idx+1, ldx, mdx)
	receiver.build(2*idx+2, mdx+1, rdx)
}

func (receiver *BookMyShow) queryMax(idx, ldx, rdx, k, maxRow int) []int {
	if ldx > maxRow {
		return []int{}
	}

	if receiver.segmentedTree[idx].max < k {
		return []int{}
	}

	if ldx == rdx {
		return []int{ldx, receiver.seats - receiver.segmentedTree[idx].max}
	}

	mdx := (ldx + rdx) / 2
	result := receiver.queryMax(2*idx+1, ldx, mdx, k, maxRow)
	if 0 != len(result) {
		return result
	}

	return receiver.queryMax(2*idx+2, mdx+1, rdx, k, maxRow)
}

func (receiver *BookMyShow) decreaseMax(idx, ldx, rdx, row, diff int) {
	if ldx > row || rdx < row {
		return
	}

	if ldx == rdx {
		receiver.segmentedTree[idx].max -= diff
		receiver.segmentedTree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	receiver.segmentedTree[idx].sum -= diff

	receiver.decreaseMax(2*idx+1, ldx, mdx, row, diff)
	receiver.decreaseMax(2*idx+2, mdx+1, rdx, row, diff)

	receiver.segmentedTree[idx].max = max2286(receiver.segmentedTree[2*idx+1].max, receiver.segmentedTree[2*idx+2].max)
}

func (receiver *BookMyShow) querySum(idx, ldx, rdx, maxRow int) int {
	if ldx > maxRow {
		return 0
	}

	if rdx <= maxRow {
		return receiver.segmentedTree[idx].sum
	}

	mdx := (ldx + rdx) / 2
	return receiver.querySum(2*idx+1, ldx, mdx, maxRow) + receiver.querySum(2*idx+2, mdx+1, rdx, maxRow)
}

func (receiver *BookMyShow) decreaseSum(idx, ldx, rdx, diff, maxRow int) {
	if ldx > maxRow {
		return
	}

	if ldx == rdx {
		receiver.segmentedTree[idx].max -= diff
		receiver.segmentedTree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	receiver.segmentedTree[idx].sum -= diff
	if mdx+1 > maxRow || receiver.segmentedTree[2*idx+1].sum >= diff {
		receiver.decreaseSum(2*idx+1, ldx, mdx, diff, maxRow)
	} else {
		diff -= receiver.segmentedTree[2*idx+1].sum

		receiver.decreaseSum(2*idx+1, ldx, mdx, receiver.segmentedTree[2*idx+1].sum, maxRow)
		receiver.decreaseSum(2*idx+2, mdx+1, rdx, diff, maxRow)
	}

	receiver.segmentedTree[idx].max = max2286(receiver.segmentedTree[2*idx+1].max, receiver.segmentedTree[2*idx+2].max)
}

func (receiver *BookMyShow) Gather(k int, maxRow int) []int {
	result := receiver.queryMax(0, 0, receiver.rows-1, k, maxRow)

	if 0 != len(result) {
		receiver.decreaseMax(0, 0, receiver.rows-1, result[0], k)
	}

	return result
}

func (receiver *BookMyShow) Scatter(k int, maxRow int) bool {
	cnt := receiver.querySum(0, 0, receiver.rows-1, maxRow)

	result := cnt >= k
	if result {
		receiver.decreaseSum(0, 0, receiver.rows-1, k, maxRow)
	}

	return result
}

func max2286(a, b int) int {
	if a > b {
		return a
	}

	return b
}
