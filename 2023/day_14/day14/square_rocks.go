package day14

type square_rocks_type struct {
	square_rocks map[int]map[int]bool
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

func (s *square_rocks_type) get_stop_pos_for_round_rock(row_index int, column_index int) int {
	new_row_index := 0
	for square_rock_row_index := range s.square_rocks[column_index] {
		if square_rock_row_index < row_index && square_rock_row_index+1 > new_row_index {
			new_row_index = square_rock_row_index + 1
		} else if square_rock_row_index == row_index {
			for {
				square_rock_row_index++
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
