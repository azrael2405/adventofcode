package day10

import (
	"fmt"
	"helper"
	"slices"
	"time"
)

func run_loop_part1(_data_array []string, _start_connections []*position_type) int {
	position_value := 0
	current_positions := _start_connections
	for {
		if current_positions[0].x == current_positions[1].x && current_positions[0].y == current_positions[1].y{
			position_value = current_positions[0].value
			break
		}
		next_positions := []*position_type{}
		for _, current_pos := range current_positions{
			current_directions := get_directions_from_position(_data_array, current_pos.x, current_pos.y)
			from_dir_index := slices.Index(current_directions, current_pos.from)
			next_dir := current_directions[(from_dir_index+1)%2]
			position_offset := go_to(next_dir)
			next_positions = append(next_positions, &position_type{
				x: current_pos.x + position_offset.x,
				y: current_pos.y + position_offset.y,
				value: current_pos.value + 1,
				from: opposite_direction(next_dir),
			})
		}

		current_positions = next_positions
	}
	return position_value
}


func run_game_part1(_data_array []string) int{
	start_pos := find_start_pos(_data_array)
	start_connections := find_start_connections(_data_array, start_pos)
	return run_loop_part1(_data_array, start_connections)
}


func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := run_game_part1(_data)
	fmt.Println("Answer 1:", answer)
	return answer
}