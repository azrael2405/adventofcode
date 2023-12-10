package day10

import (
	"regexp"
	"slices"
)


type direction_type int
const(
	LEFT direction_type = iota
	UP
	RIGHT
	DOWN
)

type position_type struct{
	x int
	y int
	from direction_type
	value int
}


var pipe_directions = map[string][]direction_type{
	"|": {UP, DOWN},
	"-": {LEFT, RIGHT},
	"L": {UP, RIGHT},
	"J": {UP, LEFT},
	"7": {DOWN, LEFT},
	"F": {DOWN, RIGHT},
	".": {},
	"S": {},
}

func go_to(_direction direction_type) *position_type{
	position_offset := &position_type{x:0, y:0}
	switch _direction{
	case UP:
		position_offset.x -= 1
	case DOWN:
		position_offset.x += 1
	case LEFT:
		position_offset.y -= 1
	case RIGHT:
		position_offset.y += 1
	}
	return position_offset
}

func opposite_direction(current_direction direction_type) direction_type{
	return (current_direction + 2) % 4
}

func find_start_pos(_data_array []string)*position_type{
	s_regex := regexp.MustCompile(`S{1}`)
	s_position := &position_type{value: 0}
	for x, data_line := range _data_array{
		positions := s_regex.FindIndex([]byte(data_line))
		if len(positions) > 0{
			s_position.x = x
			s_position.y = positions[0]
			
			break
		}
	}
	return s_position
}

func find_start_connections(_data_array []string, start_pos *position_type) []*position_type{
	positions_to_check := []*position_type{
		{
			x: start_pos.x-1,
			y: start_pos.y,
			from: DOWN,
			value: 1,
		},
		{
			x: start_pos.x+1,
			y: start_pos.y,
			from: UP,
			value: 1,
		},
		{
			x: start_pos.x,
			y: start_pos.y-1,
			from: RIGHT,
			value: 1,
		},
		{
			x: start_pos.x,
			y: start_pos.y+1,
			from: LEFT,
			value: 1,
		},
	}
	found_positions := []*position_type{}
	for _, check_pos := range positions_to_check{
		if check_pos.x < 0 || check_pos.x > len(_data_array) || check_pos.y < 0 || check_pos.y > len(_data_array[0]){
			continue
		}
		pipe_dirs := get_directions_from_position(_data_array, check_pos.x, check_pos.y)
		if len(pipe_dirs) > 0{
			if slices.Contains(pipe_dirs, check_pos.from){
				found_positions = append(found_positions, check_pos)
			}
		}
	}
	return found_positions
}


func get_directions_from_position(_data_array []string, x, y int)[]direction_type{
	pipe_symbol := string(_data_array[x][y])
	return pipe_directions[pipe_symbol]
}


func run_loop(_data_array []string, _start_connections []*position_type) int {
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


func run_game(_data_array []string) int{
	start_pos := find_start_pos(_data_array)
	start_connections := find_start_connections(_data_array, start_pos)
	return run_loop(_data_array, start_connections)
}