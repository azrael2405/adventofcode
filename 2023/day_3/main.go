package main

import (
	"fmt"
	"helper"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var x_max = 0
var y_max = 0


type position struct{
	x int
	y_start int
	y_end int
	value string
}


func check_error(e error){
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


func find_adjacent_numbers(_symbol_map map[string]bool, _number_array []*position) []int{
	return_list := []int{}
	for _, number_pos := range _number_array{
		adjacent_positions := get_adjacent_positions(number_pos)
		for _, adjacent_position := range adjacent_positions{
			if _, ok := _symbol_map[adjacent_position]; ok {
				value, conversion_error := strconv.Atoi(number_pos.value)
				check_error(conversion_error)
				return_list = append(return_list, value)
				break
			}
		}
	}
	return return_list
}


func find_gear_ratios(_symbol_map map[string]bool, _number_array []*position) []int{
	return_list := []int{}
	gears := map[string][]int{}
	for _, number_pos := range _number_array{
		adjacent_positions := get_adjacent_positions(number_pos)
		for _, adjacent_position := range adjacent_positions{
			if _, ok := _symbol_map[adjacent_position]; ok {
				value, conversion_error := strconv.Atoi(number_pos.value)
				check_error(conversion_error)
				gears[adjacent_position] = append(gears[adjacent_position], value)
			}
		}
	}
	for _, value := range gears{
		if len(value) > 1{
			ratio := 1
			for _, v := range value{
				ratio *= v
			}
			return_list = append(return_list, ratio)
		}
	}
	return return_list
}


func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	check_error(file_error)
	lines := strings.Fields(string(file_data))
	return lines
}


func parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := 0
	symbol_positions := find_symbol_positions(_data, `[^\.\d]{1}`)
	number_positions := find_number_positions(_data)
	adjacent_numbers := find_adjacent_numbers(symbol_positions, number_positions)
	for _, value := range adjacent_numbers {
		answer += value
	}
	fmt.Println("Answer 1:", answer)
}


func parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := 0
	gear_positions := find_symbol_positions(_data, `\*{1}`)
	number_positions := find_number_positions(_data)
	gear_ratios := find_gear_ratios(gear_positions, number_positions)
	for _, value := range gear_ratios {
		answer += value
	}
	fmt.Println("Answer 2:", answer)
	
}


func main (){
	defer helper.TimeTrack(time.Now(), "main")
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	x_max = len(data_array)
	y_max = len(data_array[0])
	parse_answer_one(data_array)
	parse_answer_two(data_array)
}
