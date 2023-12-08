package day8

import (
	"fmt"
	"helper"
	"time"
)

func find_exit_all(_data_map *map_type) int{
	moves := 0
	current_nodes := _data_map.nodes_with_A_ending
	moves_list := []int{}

	for {
		next_move := string(_data_map.movement_string[moves%len(_data_map.movement_string)])
		next_nodes := []*node_type{}
		moves += 1
		// fmt.Println(moves, "::", next_move)
		for _, current_node := range current_nodes{
			if next_move == "L"{
				// fmt.Println(current_node.name, current_node.left.name)
				current_node = current_node.left
			} else{
				// fmt.Println(current_node.name, current_node.right.name)
				current_node = current_node.right
			}
			if string(current_node.name[2]) == "Z"{
				moves_list = append(moves_list, moves)
			}else{
				next_nodes = append(next_nodes, current_node)
			}
		}
		// fmt.Println(moves, "::", next_move, ":", z_nodes, "/", len(next_nodes))
		if len(next_nodes) == 0{
			break
		}
		current_nodes = next_nodes
	}


		
	return LCM(moves_list)
}

func GCD(a, b int) int{
	for b != 0{
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers []int) int{
	a := integers[0]
	b := integers[1]
	result := a*b/GCD(a,b)
	if len(integers) > 2{
		for _, i:= range integers[2:] {
			result = LCM([]int{result, i})
		}
	}
	return result
}

func Parse_answer_two(_data []string)int{
	defer helper.TimeTrack(time.Now(), "Answer 2")
	data_map := parse_data(_data)
	answer := find_exit_all(data_map)
	fmt.Println("Answer 2:", answer)
	return answer
}