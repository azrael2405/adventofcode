package day14

func parse_platform_map(_data_array []string) int {
	square_rocks := &square_rocks_type{square_rocks: map[int]map[int]bool{}}
	round_rocks := &round_rocks_type{round_rocks: map[int]map[int]bool{}, rocks_weigth: 0}
	for row_index, _row := range _data_array {
		for column_index, _char := range _row {
			if string(_char) == "#" {
				square_rocks.add_square_rock(row_index, column_index)
			} else if string(_char) == "O" {
				round_rocks.add_round_rock(row_index, column_index)
			}
		}
	}
	// println("before:")
	// square_rocks.print_map(len(_data_array), len(_data_array[0]))
	round_rocks.tilt_north(square_rocks, len(_data_array))
	// println("after:")
	// square_rocks.print_map(len(_data_array), len(_data_array[0]))
	return round_rocks.rocks_weigth
}

func parse_data(_data_array []string) int {
	return parse_platform_map(_data_array)
}
