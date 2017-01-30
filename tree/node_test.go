package tree

import "testing"

func TestNode_Plus(t *testing.T) {
	type fields struct {
		Value int
		Left  *Node
		Right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Plus(); got != tt.want {
				t.Errorf("Node.Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}
