package piscine

func ListReverse(l *List) {
	curF := l.Head
	curL := l.Tail
	for curF != curL {
		curF, curL = curL, curF
		curL--
		curF++
	}
	return
}
