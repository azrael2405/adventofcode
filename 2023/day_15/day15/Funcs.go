package day15

import (
	"strings"
)

func hash_data(_data_string string) int {
	string_value := 0
	for _, char := range _data_string {
		ascii_value := int(char)
		string_value += ascii_value
		string_value *= 17
		string_value %= 256
	}
	return string_value
}

func parse_data(_data_array []string) int {
	answer_value := 0
	for _, data_string := range _data_array {
		answer_value += hash_data(data_string)
	}
	return answer_value
}

func parse_label(_data_string string) string {
	label := ""
	if strings.Contains(_data_string, "=") {
		label = _data_string[:strings.Index(_data_string, "=")]
	} else if strings.Contains(_data_string, "-") {
		label = _data_string[:strings.Index(_data_string, "-")]
	}
	return label
}

func parse_data_part_two(_data_array []string) int {
	answer_value := 0
	boxes := make(map[int]*box_type, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = &box_type{box_number: i + 1, content: make(map[string]int), positions: []string{}}
	}
	for _, data_string := range _data_array {
		label := parse_label(data_string)
		box_number := hash_data(label)
		boxes[box_number].parse_string(data_string)
	}
	for _, box := range boxes {
		answer_value += box.get_result()
	}
	return answer_value
}
