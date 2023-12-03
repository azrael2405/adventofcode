package day3

import (
	"fmt"
	"regexp"
)

var x_max int = 0
var y_max int = 0

func Set_my_globals(x, y int){
	x_max = x
	y_max = y
}

type position struct{
	x int
	y_start int
	y_end int
	value string
}

func Check_error(e error){
	if e != nil{
		panic(e)
	}
}

func create_position_string(x,y int) string{
	return fmt.Sprintf("%d;%d", x, y)
}


func get_adjacent_positions(_pos *position) []string{
	adjacent_positions := []string{}
	for y := _pos.y_start-1; y <= _pos.y_end +1 ; y += 1{
		if y < 0 || y >= y_max{
			continue
		}
		if y < _pos.y_start || y > _pos.y_end{
			adjacent_positions = append(
				adjacent_positions,
				create_position_string(_pos.x, y),
			)
		}
		
		if _pos.x - 1 >= 0{
			adjacent_positions = append(
				adjacent_positions,
				create_position_string(_pos.x-1, y),
			)	
		}
		if _pos.x + 1 < x_max{
			adjacent_positions = append(
				adjacent_positions,
				create_position_string(_pos.x+1, y),
			)	
		}
	}
	return adjacent_positions
}


func find_symbol_positions(_data_array []string, search_string string) map[string]bool  {
	re := regexp.MustCompile(search_string)
	position_map := map[string]bool{}
	for x, line := range _data_array{
		result := re.FindAllStringIndex(line, -1)
		for _, index := range result{
			new_pos := create_position_string(x,index[0])
			position_map[new_pos] = true
		}
	}
	return position_map
}


func find_number_positions(_data_array []string) []*position{
	re := regexp.MustCompile(`\d+`)
	positions_list := []*position{}
	for x, line := range _data_array{
		result := re.FindAllStringIndex(line, -1)
		for _, index := range result{
			new_pos := position{x: x, y_start: index[0], y_end: index[1]-1, value: line[index[0]:index[1]]}
			positions_list = append(positions_list, &new_pos)
		}
	}
	return positions_list
}