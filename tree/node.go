package tree

// Node struct
type Node struct {
	Value       int
	Left, Right *Node
}

// Plus doubles the value
func (n Node) Plus() int {
	return 2 * n.Value
}
