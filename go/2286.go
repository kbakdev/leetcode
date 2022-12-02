package _go

type BookMyShow struct {
	rows, seats int
	stree       []segNode2286
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

func (this *BookMyShow) build(idx, ldx, rdx int) {
	if ldx == rdx {
		this.stree[idx] = segNode2286{this.seats, this.seats}
		return
	}

	mdx := (ldx + rdx) / 2
	this.stree[idx] = segNode2286{this.seats, (rdx - ldx + 1) * this.seats}

	this.build(2*idx+1, ldx, mdx)
	this.build(2*idx+2, mdx+1, rdx)
}

func (this *BookMyShow) queryMax(idx, ldx, rdx, k, maxRow int) []int {
	if ldx > maxRow {
		return []int{}
	}

	if this.stree[idx].max < k {
		return []int{}
	}

	if ldx == rdx {
		return []int{ldx, this.seats - this.stree[idx].max}
	}

	mdx := (ldx + rdx) / 2
	result := this.queryMax(2*idx+1, ldx, mdx, k, maxRow)
	if 0 != len(result) {
		return result
	}

	return this.queryMax(2*idx+2, mdx+1, rdx, k, maxRow)
}

func (this *BookMyShow) decreaseMax(idx, ldx, rdx, row, diff int) {
	if ldx > row || rdx < row {
		return
	}

	if ldx == rdx {
		this.stree[idx].max -= diff
		this.stree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	this.stree[idx].sum -= diff

	this.decreaseMax(2*idx+1, ldx, mdx, row, diff)
	this.decreaseMax(2*idx+2, mdx+1, rdx, row, diff)

	this.stree[idx].max = max2286(this.stree[2*idx+1].max, this.stree[2*idx+2].max)
}

func (this *BookMyShow) querySum(idx, ldx, rdx, maxRow int) int {
	if ldx > maxRow {
		return 0
	}

	if rdx <= maxRow {
		return this.stree[idx].sum
	}

	mdx := (ldx + rdx) / 2
	return this.querySum(2*idx+1, ldx, mdx, maxRow) + this.querySum(2*idx+2, mdx+1, rdx, maxRow)
}

func (this *BookMyShow) decreaseSum(idx, ldx, rdx, diff, maxRow int) {
	if ldx > maxRow {
		return
	}

	if ldx == rdx {
		this.stree[idx].max -= diff
		this.stree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	this.stree[idx].sum -= diff
	if mdx+1 > maxRow || this.stree[2*idx+1].sum >= diff {
		this.decreaseSum(2*idx+1, ldx, mdx, diff, maxRow)
	} else {
		diff -= this.stree[2*idx+1].sum

		this.decreaseSum(2*idx+1, ldx, mdx, this.stree[2*idx+1].sum, maxRow)
		this.decreaseSum(2*idx+2, mdx+1, rdx, diff, maxRow)
	}

	this.stree[idx].max = max2286(this.stree[2*idx+1].max, this.stree[2*idx+2].max)
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	result := this.queryMax(0, 0, this.rows-1, k, maxRow)

	if 0 != len(result) {
		this.decreaseMax(0, 0, this.rows-1, result[0], k)
	}

	return result
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	cnt := this.querySum(0, 0, this.rows-1, maxRow)

	result := cnt >= k
	if result {
		this.decreaseSum(0, 0, this.rows-1, k, maxRow)
	}

	return result
}

func max2286(a, b int) int {
	if a > b {
		return a
	}

	return b
}
