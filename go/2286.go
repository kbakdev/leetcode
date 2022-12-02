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

func (r *BookMyShow) build(idx, ldx, rdx int) {
	if ldx == rdx {
		r.stree[idx] = segNode2286{r.seats, r.seats}
		return
	}

	mdx := (ldx + rdx) / 2
	r.stree[idx] = segNode2286{r.seats, (rdx - ldx + 1) * r.seats}

	r.build(2*idx+1, ldx, mdx)
	r.build(2*idx+2, mdx+1, rdx)
}

func (r *BookMyShow) queryMax(idx, ldx, rdx, k, maxRow int) []int {
	if ldx > maxRow {
		return []int{}
	}

	if r.stree[idx].max < k {
		return []int{}
	}

	if ldx == rdx {
		return []int{ldx, r.seats - r.stree[idx].max}
	}

	mdx := (ldx + rdx) / 2
	result := r.queryMax(2*idx+1, ldx, mdx, k, maxRow)
	if 0 != len(result) {
		return result
	}

	return r.queryMax(2*idx+2, mdx+1, rdx, k, maxRow)
}

func (r *BookMyShow) decreaseMax(idx, ldx, rdx, row, diff int) {
	if ldx > row || rdx < row {
		return
	}

	if ldx == rdx {
		r.stree[idx].max -= diff
		r.stree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	r.stree[idx].sum -= diff

	r.decreaseMax(2*idx+1, ldx, mdx, row, diff)
	r.decreaseMax(2*idx+2, mdx+1, rdx, row, diff)

	r.stree[idx].max = max2286(r.stree[2*idx+1].max, r.stree[2*idx+2].max)
}

func (r *BookMyShow) querySum(idx, ldx, rdx, maxRow int) int {
	if ldx > maxRow {
		return 0
	}

	if rdx <= maxRow {
		return r.stree[idx].sum
	}

	mdx := (ldx + rdx) / 2
	return r.querySum(2*idx+1, ldx, mdx, maxRow) + r.querySum(2*idx+2, mdx+1, rdx, maxRow)
}

func (r *BookMyShow) decreaseSum(idx, ldx, rdx, diff, maxRow int) {
	if ldx > maxRow {
		return
	}

	if ldx == rdx {
		r.stree[idx].max -= diff
		r.stree[idx].sum -= diff

		return
	}

	mdx := (ldx + rdx) / 2
	r.stree[idx].sum -= diff
	if mdx+1 > maxRow || r.stree[2*idx+1].sum >= diff {
		r.decreaseSum(2*idx+1, ldx, mdx, diff, maxRow)
	} else {
		diff -= r.stree[2*idx+1].sum

		r.decreaseSum(2*idx+1, ldx, mdx, r.stree[2*idx+1].sum, maxRow)
		r.decreaseSum(2*idx+2, mdx+1, rdx, diff, maxRow)
	}

	r.stree[idx].max = max2286(r.stree[2*idx+1].max, r.stree[2*idx+2].max)
}

func (r *BookMyShow) Gather(k int, maxRow int) []int {
	result := r.queryMax(0, 0, r.rows-1, k, maxRow)

	if 0 != len(result) {
		r.decreaseMax(0, 0, r.rows-1, result[0], k)
	}

	return result
}

func (r *BookMyShow) Scatter(k int, maxRow int) bool {
	cnt := r.querySum(0, 0, r.rows-1, maxRow)

	result := cnt >= k
	if result {
		r.decreaseSum(0, 0, r.rows-1, k, maxRow)
	}

	return result
}

func max2286(a, b int) int {
	if a > b {
		return a
	}

	return b
}
