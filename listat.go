package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	i := 0

	for current != nil {
		if i == pos {
			return current
		}
		i++
		current = current.Next
	}
	return nil
}
