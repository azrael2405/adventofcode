package day14

import "math"

type square_rocks_type struct {
	square_rocks            map[int]map[int]bool
	square_rocks_transposed map[int]map[int]bool
	original_square_rocks   map[int]map[int]bool
	max_ranges              []int
}

func (s *square_rocks_type) reset(original map[int]map[int]bool) { //square_rocks_type) {
	println()
	for i := 0; i < s.max_ranges[1]; i++ {
		for j := 0; j < s.max_ranges[3]; j++ {
			if _, ok := original[j][i]; ok {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
	s.square_rocks = map[int]map[int]bool{}
	for k, v := range original {
		x := map[int]bool{}
		for k2, v2 := range v {
			x[k2] = v2
		}
		s.square_rocks[k] = x
	}
	s.square_rocks_transposed = transpose(s.square_rocks)
}

func (s *square_rocks_type) print_map(_max_rows, _max_columns int) {
	for row_index := 0; row_index < _max_rows; row_index++ {
		for column_index := 0; column_index < _max_columns; column_index++ {
			if _, ok := s.square_rocks[column_index][row_index]; ok {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func (s *square_rocks_type) add_square_rock(row_index int, column_index int) {
	if _, ok := s.square_rocks[column_index]; !ok {
		s.square_rocks[column_index] = map[int]bool{}
	}
	s.square_rocks[column_index][row_index] = true
}

func (s *square_rocks_type) add_square_rock_transposed(row_index int, column_index int) {
	if _, ok := s.square_rocks_transposed[column_index]; !ok {
		s.square_rocks_transposed[column_index] = map[int]bool{}
	}
	s.square_rocks_transposed[column_index][row_index] = true
}

func (s *square_rocks_type) get_stop_pos_north(row_index int, column_index int) int {
	new_row_index := 0
	for square_rock_row_index := range s.square_rocks[column_index] {
		if square_rock_row_index < row_index && square_rock_row_index+1 > new_row_index {
			new_row_index = square_rock_row_index + 1
		} else if square_rock_row_index == row_index {
			for {
				square_rock_row_index++
				if square_rock_row_index >= s.max_ranges[1] {
					square_rock_row_index--
					break
				}
				if _, ok := s.square_rocks[column_index][square_rock_row_index]; !ok {
					break
				}
			}
			if new_row_index < square_rock_row_index {
				new_row_index = square_rock_row_index
			}
		}
	}
	s.add_square_rock(new_row_index, column_index)
	return new_row_index
}

func (s *square_rocks_type) get_stop_pos_south(row_index int, column_index int) int {
	new_row_index := math.MaxInt
	for square_rock_row_index := range s.square_rocks[column_index] {
		if square_rock_row_index > row_index && square_rock_row_index-1 < new_row_index {
			new_row_index = square_rock_row_index - 1
		} else if square_rock_row_index == row_index {
			for {
				square_rock_row_index--
				if square_rock_row_index < s.max_ranges[0] {
					square_rock_row_index++
					break
				}
				if _, ok := s.square_rocks[column_index][square_rock_row_index]; !ok {
					break
				}
			}
			if new_row_index > square_rock_row_index {
				new_row_index = square_rock_row_index
			}
		}
	}
	s.add_square_rock(new_row_index, column_index)
	return new_row_index
}

func (s *square_rocks_type) get_stop_pos_east(row_index int, column_index int) int {
	new_column_index := 0
	for square_rock_column_index := range s.square_rocks_transposed[row_index] {
		if square_rock_column_index < column_index && square_rock_column_index+1 > new_column_index {
			new_column_index = square_rock_column_index + 1
		} else if square_rock_column_index == column_index {
			for {
				square_rock_column_index++
				if square_rock_column_index >= s.max_ranges[3] {
					square_rock_column_index--
					break
				}
				if _, ok := s.square_rocks_transposed[row_index][square_rock_column_index]; !ok {
					break
				}
			}
			if new_column_index < square_rock_column_index {
				new_column_index = square_rock_column_index
			}
		}
	}
	s.add_square_rock_transposed(new_column_index, row_index)
	return new_column_index
}

func (s *square_rocks_type) get_stop_pos_west(row_index int, column_index int) int {
	new_column_index := math.MaxInt
	for square_rock_column_index := range s.square_rocks_transposed[row_index] {
		if square_rock_column_index > column_index && square_rock_column_index+1 < new_column_index {
			new_column_index = square_rock_column_index - 1
		} else if square_rock_column_index == column_index {
			for {
				square_rock_column_index--
				if square_rock_column_index < s.max_ranges[2] {
					square_rock_column_index++
					break
				}
				if _, ok := s.square_rocks_transposed[row_index][square_rock_column_index]; !ok {
					break
				}
			}
			if new_column_index > square_rock_column_index {
				new_column_index = square_rock_column_index
			}
		}
	}
	s.add_square_rock_transposed(new_column_index, row_index)
	return new_column_index
}
