package structs

type Move struct {
	Direction Direction
	Steps     int
	From      *Node
	To        *Node
	Score     int
}

func (m *Move) create_new(to *Node, dir Direction) *Move {
	steps := 0
	if dir == m.Direction {
		steps = m.Steps + 1
	}
	if steps > 2 {
		return nil
	}
	return &Move{Direction: dir, Steps: steps, From: m.To, To: to, Score: m.Score + to.Base_value}
}

func (m *Move) Get_next_moves(nodeMap *NodeMap) []*Move {
	moves := make([]*Move, 0)
	for _, dir := range []Direction{UP, RIGHT, DOWN, LEFT} {
		next_node := nodeMap.Get_node(m.To.Position, dir)
		if next_node == nil || next_node == m.From {
			continue
		}
		new_move := m.create_new(next_node, dir)
		if new_move == nil {
			continue
		}
		moves = append(moves, new_move)
	}
	return moves
}
