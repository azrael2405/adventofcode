package day10

import (
	"fmt"
	"helper"
	"slices"
	"time"
)

func find_right_side(from, to direction_type) bool{
	switch {
	case from == DOWN && (to == UP || to == RIGHT):
		return false
	case from == RIGHT && to == UP:
		return false
	default:
		return true
	}
}


func run_loop_part2(_data_array []string, _start_connection *position_type) int {
	loop_positions := []*position_type{}
	inside_positions := 0
	current_position := _start_connection
	
	for {
		current_directions := get_directions_from_position(_data_array, current_position.x, current_position.y)
		from_dir_index := slices.Index(current_directions, current_position.from)
		next_dir := current_directions[(from_dir_index+1)%2]
		current_position.to = next_dir
		if slices.ContainsFunc(loop_positions, func (element *position_type) bool{
			if element.x == current_position.x && element.y == current_position.y{
				return true
			}
			return false
		} ){
			break
		}
		loop_positions = append(loop_positions, current_position)
		position_offset := go_to(next_dir)
		next_position := &position_type{
			x: current_position.x + position_offset.x,
			y: current_position.y + position_offset.y,
			from: opposite_direction(next_dir),
		}
		current_position = next_position
	}
	slices.SortFunc(loop_positions, func(a, b *position_type)int{
		value := 0
		switch{
		case a.x < b.x:
			value = -1
		case a.x > b.x:
			value = 1
		case a.y < b.y:
			value = -1
		case a.y > b.y:
			value = 1
		}
		return value
	})
	for x, data_line := range _data_array{
		for y := range data_line{
			if slices.ContainsFunc(loop_positions, func (element *position_type) bool{
				if element.x == x && element.y == y{
					return true
				}
				return false
			}){
				continue
			}
			next_loop_index := slices.IndexFunc(loop_positions, func(element *position_type)bool{
				if element.x == x && element.y > y{
					return true
				}
				return false
			})
			if next_loop_index != -1{
				if! find_right_side(loop_positions[next_loop_index].from, loop_positions[next_loop_index].to){
					inside_positions += 1
					
				}
			}
		}
	}
	return inside_positions
}


func run_game_part2(_data_array []string) int{
	start_pos := find_start_pos(_data_array)
	start_connections := find_start_connections(_data_array, start_pos)
	pipe_directions["S"] = []direction_type{
		opposite_direction(start_connections[0].from),
		opposite_direction(start_connections[1].from),
	}
	fmt.Println(pipe_directions["S"])
	return run_loop_part2(_data_array, start_pos)
}


func Parse_answer_two(_data []string)int{
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := run_game_part2(_data)
	fmt.Println("Answer 2:", answer)
	return answer
}