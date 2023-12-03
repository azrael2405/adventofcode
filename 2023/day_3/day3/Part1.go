package day3

import (
	"fmt"
	"helper"
	"strconv"
	"time"
)


func find_adjacent_numbers(_symbol_map map[string]bool, _number_array []*position) []int{
	return_list := []int{}
	for _, number_pos := range _number_array{
		adjacent_positions := get_adjacent_positions(number_pos)
		for _, adjacent_position := range adjacent_positions{
			if _, ok := _symbol_map[adjacent_position]; ok {
				value, conversion_error := strconv.Atoi(number_pos.value)
				Check_error(conversion_error)
				return_list = append(return_list, value)
				break
			}
		}
	}
	return return_list
}


func Parse_answer_one(_data []string){
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