package structs

import "math"

type NodeMap struct {
	Nodes [][]Node
	Costs map[*Node]int
}

func (n *NodeMap) Init(_data_array []string) {
	n.Nodes = make([][]Node, len(_data_array))
	n.Costs = make(map[*Node]int)
	for y, line := range _data_array {
		n.Nodes[y] = make([]Node, len(line))
		for x, value := range line {
			n.Nodes[y][x].Init(x, y, string(value))
			n.Costs[&n.Nodes[y][x]] = math.MaxInt
		}
	}
	n.Costs[&n.Nodes[0][0]] = 0
}

func (n *NodeMap) Get_node(_current_position Position, direction Direction) *Node {
	switch direction {
	case UP:
		if _current_position.Y-1 < 0 || _current_position.Y-1 >= len(n.Nodes) {
			return nil
		}
		return &n.Nodes[_current_position.Y-1][_current_position.X]
	case RIGHT:
		if _current_position.X+1 < 0 || _current_position.X+1 >= len(n.Nodes[0]) {
			return nil
		}
		return &n.Nodes[_current_position.Y][_current_position.X+1]
	case DOWN:
		if _current_position.Y+1 < 0 || _current_position.Y+1 >= len(n.Nodes) {
			return nil
		}
		return &n.Nodes[_current_position.Y+1][_current_position.X]
	case LEFT:
		if _current_position.X-1 < 0 || _current_position.X-1 >= len(n.Nodes[0]) {
			return nil
		}
		return &n.Nodes[_current_position.Y][_current_position.X-1]
	}
	return nil
}
