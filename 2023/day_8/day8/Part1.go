package day8

import (
	"fmt"
	"helper"
	"time"
)

func find_exit(_data_map *map_type) int{
	exit_name := "ZZZ"
	moves := 0
	current_node := _data_map.head
	for current_node.name != exit_name{
		next_move := string(_data_map.movement_string[moves%len(_data_map.movement_string)])
		// fmt.Println(current_node.name, next_move)
		moves += 1
		if next_move == "L"{
			current_node = current_node.left
		} else{
			current_node = current_node.right
		}
	}
	return moves
}

func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	data_map := parse_data(_data)
	answer := find_exit(data_map)
	fmt.Println("Answer 1:", answer)
	return 0
}