package day17

import (
	"day17/day17/structs"
	"strings"
)

func print_move_map(_move_map map[*structs.Node]*structs.Move, _data_array []string, _node_map *structs.NodeMap) {
	// start_position := structs.Position{X: 0, Y: 0}
	end_position := structs.Position{X: len(_data_array[0]) - 1, Y: len(_data_array) - 1}
	end_node := &_node_map.Nodes[end_position.Y][end_position.X]
	array_copy := make([]string, len(_data_array))
	copy(array_copy, _data_array)
	symbol := " "

	for end_node != nil {
		// println(end_node.Position.X, end_node.Position.Y)
		switch _move_map[end_node].Direction {
		case structs.UP:
			symbol = "^"
		case structs.RIGHT:
			symbol = ">"
		case structs.DOWN:
			symbol = "v"
		case structs.LEFT:
			symbol = "<"
		}
		array_copy[end_node.Position.Y] = array_copy[end_node.Position.Y][:end_node.Position.X] + symbol + array_copy[end_node.Position.Y][end_node.Position.X+1:]
		end_node = _move_map[end_node].From
	}
	println()
	for y, line := range array_copy {
		for x, value := range line {
			if strings.Contains("123456789", string(value)) {

				if _, ok := _move_map[&_node_map.Nodes[y][x]]; ok {
					print(".")
				} else {
					print(string(value))
				}
			} else {
				print(string(value))
			}
		}
		println()
	}
	println()

}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func heuristic(_start_position structs.Position, _end_position structs.Position) int {
	return abs(_end_position.X - _start_position.X + _end_position.Y - _start_position.Y)
}

func get_lowest_score_index(_moves []*structs.Move) int {
	lowest_move_index := 0
	lowest_move := _moves[lowest_move_index]
	for index, move := range _moves {
		if move.Score > lowest_move.Score {
			lowest_move_index = index
			lowest_move = move
		}
	}
	return lowest_move_index
}
func sum_movement(_came_from_map map[*structs.Node]*structs.Move, end_Node *structs.Node) int {
	sum := 0
	current_node := end_Node
	for current_node != nil {
		next_node := _came_from_map[current_node].From
		if next_node == nil {
			break
		}
		sum += current_node.Base_value
		current_node = next_node
	}
	return sum
}

func find_path(_data_array []string) int {
	start_position := structs.Position{X: 0, Y: 0}
	end_position := structs.Position{X: len(_data_array[0]) - 1, Y: len(_data_array) - 1}
	open_list := make([]*structs.Move, 0)
	nodeMap := structs.NodeMap{}
	nodeMap.Init(_data_array)
	open_list = append(
		open_list,
		&structs.Move{From: nil, To: &nodeMap.Nodes[start_position.Y][start_position.X], Direction: structs.LEFT, Steps: 0, Score: 0},
		// &structs.Move{From: &nodeMap.Nodes[start_position.Y][start_position.X], To: &nodeMap.Nodes[start_position.Y][start_position.X+1], Direction: structs.RIGHT, Steps: 0, Score: nodeMap.Nodes[start_position.Y][start_position.X+1].Base_value},
		// &structs.Move{From: &nodeMap.Nodes[start_position.Y][start_position.X], To: &nodeMap.Nodes[start_position.Y+1][start_position.X], Direction: structs.DOWN, Steps: 0, Score: nodeMap.Nodes[start_position.Y+1][start_position.X].Base_value},
	)
	// visited_nodes := make(map[*structs.Node]bool)
	came_from_map := make(map[*structs.Node]*structs.Move)
	came_from_map[&nodeMap.Nodes[start_position.Y][start_position.X]] = open_list[0]
	// came_from_map[&nodeMap.Nodes[start_position.Y][start_position.X+1]] = open_list[0]
	// came_from_map[&nodeMap.Nodes[start_position.Y+1][start_position.X]] = open_list[1]
	for len(open_list) > 0 {
		move_index := get_lowest_score_index(open_list)
		current_move := open_list[move_index]
		// println(current_move.To.Position.X, current_move.To.Position.Y, current_move.Score, len(open_list))
		open_list = append(open_list[:move_index], open_list[move_index+1:]...)
		if current_move.To.Position.Equals(end_position) {
			continue
		}
		for _, next_move := range current_move.Get_next_moves(&nodeMap) {
			// println(next_move.To.Position.X, next_move.To.Position.Y, next_move.Score, len(open_list))
			if next_move.Score < nodeMap.Costs[next_move.To] {
				came_from_map[next_move.To] = next_move
				nodeMap.Costs[next_move.To] = next_move.Score
				// if _, ok := visited_nodes[next_move.To]; !ok {
				open_list = append(open_list, next_move)
				// visited_nodes[next_move.To] = true
				// }
			}
		}
	}
	print_move_map(came_from_map, _data_array, &nodeMap)
	return sum_movement(came_from_map, &nodeMap.Nodes[end_position.Y][end_position.X])
}

func parse_data(_data_array []string) int {
	answer_value := find_path(_data_array)
	return answer_value
}
