package piscine

func ListReverse(l *List) {
	var prev *NodeL // starts as nil (will become the new tail's Next)
	current := l.Head
	l.Tail = l.Head // old head becomes the new tail

	for current != nil {
		next := current.Next // save where we're going
		current.Next = prev  // flip the pointer backwards
		prev = current       // move prev forward
		current = next       // move current forward
	}
	l.Head = prev // prev is now the last node we visited = new head
}
