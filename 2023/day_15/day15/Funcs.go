package day15

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
