package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	current := l.Head
	last := current
	for current != nil {
		last = current
		current = current.Next
	}
	return last
}
