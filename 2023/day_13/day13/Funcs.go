package day13

import (
	"strings"
)

type Tuple struct {
	is_mirror   bool
	differences int
}

func count_differences(a, b string) int {
	difference_count := 0
	for index, char := range a {
		if string(char) != string(b[index]) {
			difference_count++
		}
	}
	return difference_count
}

func is_mirror_line1(_data_lines []string, _row_index int) bool {
	is_mirror := true
	for offset := 1; offset+_row_index+1 < len(_data_lines) && _row_index-offset >= 0; offset++ {
		if _data_lines[_row_index+1+offset] != _data_lines[_row_index-offset] {
			is_mirror = false
			break
		}
	}
	return is_mirror
}

func is_mirror_line2(_data_lines []string, _row_index int, differences int) *Tuple {
	is_mirror := true
	for offset := 1; offset+_row_index+1 < len(_data_lines) && _row_index-offset >= 0; offset++ {
		differences += count_differences(_data_lines[_row_index+1+offset], _data_lines[_row_index-offset])
		if differences > 1 {
			is_mirror = false
			break
		}
	}
	return &Tuple{is_mirror: is_mirror, differences: differences}
}

func parse_rows1(_data_lines []string) int {
	mirror_line := 0
	for row_index, row := range _data_lines {
		if row_index >= len(_data_lines)-1 {
			break
		}

		if row == _data_lines[row_index+1] {
			if is_mirror_line1(_data_lines, row_index) {
				mirror_line = row_index + 1
				break
			}
		}
	}
	return mirror_line
}

func parse_rows2(_data_lines []string) int {
	mirror_line := 0
	for row_index, row := range _data_lines {
		if row_index >= len(_data_lines)-1 {
			break
		}
		differences := count_differences(row, _data_lines[row_index+1])
		if differences <= 1 {
			tuple := is_mirror_line2(_data_lines, row_index, differences)
			if tuple.is_mirror && tuple.differences == 1 {
				mirror_line = row_index + 1
				break
			}
		}
	}
	return mirror_line
}

func transpose_array(_data_lines []string) []string {
	output_array := make([]string, len(_data_lines[0]))
	for _, row := range _data_lines {
		for row_position, char := range row {
			output_array[row_position] += string(char)
		}
	}
	return output_array
}

func parse_data1(_data_array []string) int {
	answer_value := 0
	for _, data_block := range _data_array {
		data := strings.Split(data_block, "\n")
		row_value := parse_rows1(data)
		transposed_data := transpose_array(data)
		column_value := parse_rows1(transposed_data)
		answer_value += (100 * row_value) + column_value

	}
	return answer_value
}

func parse_data2(_data_array []string) int {
	answer_value := 0
	for _, data_block := range _data_array {
		data := strings.Split(data_block, "\n")
		row_value := parse_rows2(data)
		transposed_data := transpose_array(data)
		column_value := parse_rows2(transposed_data)
		answer_value += (100 * row_value) + column_value

	}
	return answer_value
}
