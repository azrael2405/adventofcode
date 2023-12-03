package day3

import (
	"fmt"
	"helper"
	"strconv"
	"time"
)


func find_gear_ratios(_symbol_map map[string]bool, _number_array []*position) []int{
	return_list := []int{}
	gears := map[string][]int{}
	for _, number_pos := range _number_array{
		adjacent_positions := get_adjacent_positions(number_pos)
		for _, adjacent_position := range adjacent_positions{
			if _, ok := _symbol_map[adjacent_position]; ok {
				value, conversion_error := strconv.Atoi(number_pos.value)
				Check_error(conversion_error)
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


func Parse_answer_two(_data []string){
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