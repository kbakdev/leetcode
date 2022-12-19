package _go

func minPartitions(n string) (answer int) {
	for _, character := range n {
		if t := int(character - '0'); t > answer {
			answer = t
		}
	}
	return
}
