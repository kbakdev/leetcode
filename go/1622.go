package _go

// Fancy initializes the object with an empty sequence.
type Fancy struct {
	list []int64
}

func Constructor() Fancy {
	var receiver Fancy
	receiver.list = make([]int64, 0, 1000)
	return receiver
}

func (receiver *Fancy) Append(val int) {
	receiver.list = append(receiver.list, int64(val))
}

func (receiver *Fancy) AddAll(inc int) {
	for ix, val := range receiver.list {
		receiver.list[ix] = val + int64(inc)
	}
}

func (receiver *Fancy) MultAll(m int) {
	for ix, val := range receiver.list {
		receiver.list[ix] = val * int64(m) % 1000000007
	}
}

func (receiver *Fancy) GetIndex(idx int) int {
	var intval int
	if idx < len(receiver.list) {
		intval = int(receiver.list[idx])
	} else {
		intval = -1
	}

	return intval
}
