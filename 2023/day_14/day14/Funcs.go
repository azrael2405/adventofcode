package day14

import "fmt"

func transpose(matrix map[int]map[int]bool) map[int]map[int]bool {
	new_matrix := map[int]map[int]bool{}
	for row_index := range matrix {
		for column_index := range matrix[row_index] {
			if _, ok := new_matrix[column_index]; !ok {
				new_matrix[column_index] = map[int]bool{}
			}
			new_matrix[column_index][row_index] = true
		}
	}
	return new_matrix
}

func print_map(s *square_rocks_type, r *round_rocks_type, _max_rows, _max_columns int) {
	for row_index := 0; row_index < _max_rows; row_index++ {
		for column_index := 0; column_index < _max_columns; column_index++ {
			if _, ok := s.square_rocks[column_index][row_index]; ok {
				print("#")
			} else if _, ok := r.round_rocks[column_index][row_index]; ok {
				print("O")
			} else {
				print(".")
			}
		}
		println()
	}
}

func parse_platform_map(_data_array []string, part int) int {
	square_rocks := square_rocks_type{
		square_rocks: map[int]map[int]bool{},
		max_ranges: []int{
			0,
			len(_data_array),
			0,
			len(_data_array[0]),
		},
	}
	round_rocks := round_rocks_type{round_rocks: map[int]map[int]bool{}, rocks_weigth: 0}
	for row_index, _row := range _data_array {
		for column_index, _char := range _row {
			if string(_char) == "#" {
				square_rocks.add_square_rock(row_index, column_index)
			} else if string(_char) == "O" {
				round_rocks.add_round_rock(row_index, column_index)
			}
		}
	}

	if part == 1 {
		// println("before:")
		// square_rocks.print_map(len(_data_array), len(_data_array[0]))
		round_rocks.tilt_north(&square_rocks, len(_data_array))
		// println("after:")
		// square_rocks.print_map(len(_data_array), len(_data_array[0]))
		return round_rocks.get_weight_north(len(_data_array))
	} else {

		original_square_rocks := map[int]map[int]bool{}
		for k, v := range square_rocks.square_rocks {
			x := map[int]bool{}
			for k2, v2 := range v {
				x[k2] = v2
			}
			original_square_rocks[k] = x
		}

		square_rocks.square_rocks_transposed = transpose(square_rocks.square_rocks)
		for i := 0; i < 1; i++ { //1_000_000_000
			round_rocks.tilt_north(&square_rocks, len(_data_array))
			// original_square_rocks.print_map(len(_data_array), len(_data_array[0]))
			square_rocks.reset(original_square_rocks)
			fmt.Println("\n", i, ": north")
			fmt.Println("--------------------")
			print_map(&square_rocks, &round_rocks, len(_data_array), len(_data_array[0]))
			round_rocks.tilt_west(&square_rocks, len(_data_array[0]))
			square_rocks.reset(original_square_rocks)
			fmt.Println("\n", i, ": west")
			fmt.Println("--------------------")
			print_map(&square_rocks, &round_rocks, len(_data_array), len(_data_array[0]))
			round_rocks.tilt_south(&square_rocks, len(_data_array))
			square_rocks.reset(original_square_rocks)
			fmt.Println("\n", i, ": south")
			fmt.Println("--------------------")
			print_map(&square_rocks, &round_rocks, len(_data_array), len(_data_array[0]))
			round_rocks.tilt_east(&square_rocks, len(_data_array[0]))
			square_rocks.reset(original_square_rocks)
			fmt.Println("\n", i, ": east")
			fmt.Println("--------------------")
			print_map(&square_rocks, &round_rocks, len(_data_array), len(_data_array[0]))

			fmt.Println(i, "; Weight: ", round_rocks.get_weight_north(len(_data_array)))
		}
		return round_rocks.get_weight_north(len(_data_array))
	}
}

func parse_data(_data_array []string, part int) int {
	return parse_platform_map(_data_array, part)
}
