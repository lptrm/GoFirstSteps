package main

type vertex struct {
	value      int
	leftChild  *vertex
	rightChild *vertex
}
type tree struct {
	vertices []*vertex
}

func (t tree) insert(value int) bool {
	newVertex := &vertex{value: value}
	if len(t.vertices) == 0 {
		t.vertices = []*vertex{newVertex}
	} else {
		for _, vertex := range t.vertices {
			if vertex.value == value {
				return false
			}
		}
	}

}
func main() {
	t := tree{}
	t.insert(3)
	t.insert(2)
	t.insert(1)
}
