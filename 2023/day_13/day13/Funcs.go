package day13

import (
	"fmt"
	"strings"
)


func is_mirror_line(_data_lines []string, _row_index int) bool{
	is_mirror := true
	for offset := 1; offset + _row_index + 1 < len(_data_lines) && _row_index - offset >= 0; offset++{
		if _data_lines[_row_index + 1 + offset] != _data_lines[_row_index - offset]{
			is_mirror = false
			break
		}
	}
	return is_mirror
}

func parse_rows(_data_lines []string) int{
	mirror_line := 0
	for row_index, row := range _data_lines{
		if row_index >= len(_data_lines)-1{
			break
		}
		if row == _data_lines[row_index+1]{
			if is_mirror_line(_data_lines, row_index){
				mirror_line = row_index + 1
				break
			}
		}
	}
	fmt.Println(mirror_line)
	return mirror_line
}

func transpose_array(_data_lines []string)[]string{
	output_array := make([]string, len(_data_lines[0]))
	for _, row := range _data_lines{
		for row_position, char := range row{
			output_array[row_position] += string(char)
		}
	} 
	return output_array
}


func parse_data(_data_array []string) int{
	answer_value := 0
	for _, data_block := range _data_array{
		data := strings.Split(data_block, "\n")
		transposed_data := transpose_array(data)
		column_value := parse_rows(transposed_data)
		row_value := parse_rows(data)
		answer_value += (100 * row_value) + column_value
		
	}
	return answer_value
}