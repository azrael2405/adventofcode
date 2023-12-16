package day14

type round_rocks_type struct {
	round_rocks            map[int]map[int]bool
	round_rocks_transposed map[int]map[int]bool
	rocks_weigth           int
}

func (r *round_rocks_type) print_map(_max_rows, _max_columns int) {
	for row_index := 0; row_index < _max_rows; row_index++ {
		for column_index := 0; column_index < _max_columns; column_index++ {
			if _, ok := r.round_rocks[column_index][row_index]; ok {
				print("O")
			} else {
				print(".")
			}
		}
		println()
	}
}

func (r *round_rocks_type) add_round_rock(row_index int, column_index int) {
	if _, ok := r.round_rocks[column_index]; !ok {
		r.round_rocks[column_index] = map[int]bool{}
	}
	r.round_rocks[column_index][row_index] = true
}

func (r *round_rocks_type) tilt_north(s *square_rocks_type, _max_rows int) {
	new_round_rocks := map[int]map[int]bool{}
	r.rocks_weigth = 0
	for column_index := range r.round_rocks {
		new_rock_column := map[int]bool{}
		for row_index := range r.round_rocks[column_index] {
			new_column_index := s.get_stop_pos_north(row_index, column_index)
			new_rock_column[new_column_index] = true

		}
		new_round_rocks[column_index] = new_rock_column
	}
	r.round_rocks = new_round_rocks
	r.round_rocks_transposed = transpose(r.round_rocks)
}

func (r *round_rocks_type) tilt_south(s *square_rocks_type, _max_rows int) {
	new_round_rocks := map[int]map[int]bool{}
	r.rocks_weigth = 0
	for column_index := range r.round_rocks {
		new_rock_column := map[int]bool{}
		for row_index := range r.round_rocks[column_index] {
			new_column_index := s.get_stop_pos_south(row_index, column_index)
			new_rock_column[new_column_index] = true

		}
		new_round_rocks[column_index] = new_rock_column
	}
	r.round_rocks = new_round_rocks
	r.round_rocks_transposed = transpose(r.round_rocks)
}

func (r *round_rocks_type) tilt_west(s *square_rocks_type, _max_cols int) {
	new_round_rocks := map[int]map[int]bool{}
	r.rocks_weigth = 0
	for row_index := range r.round_rocks_transposed {
		new_rock_row := map[int]bool{}
		for column_index := range r.round_rocks_transposed[row_index] {
			new_row_index := s.get_stop_pos_west(row_index, column_index)
			new_rock_row[new_row_index] = true

		}
		new_round_rocks[row_index] = new_rock_row
	}
	r.round_rocks_transposed = new_round_rocks
	r.round_rocks = transpose(r.round_rocks_transposed)
}

func (r *round_rocks_type) tilt_east(s *square_rocks_type, _max_cols int) {
	new_round_rocks := map[int]map[int]bool{}
	r.rocks_weigth = 0
	for row_index := range r.round_rocks_transposed {
		new_rock_row := map[int]bool{}
		for column_index := range r.round_rocks_transposed[row_index] {
			new_row_index := s.get_stop_pos_south(row_index, column_index)
			new_rock_row[new_row_index] = true

		}
		new_round_rocks[row_index] = new_rock_row
	}
	r.round_rocks_transposed = new_round_rocks
	r.round_rocks = transpose(r.round_rocks_transposed)
}

func (r *round_rocks_type) get_weight_north(max_rows int) int {
	r.rocks_weigth = 0
	for column_index := range r.round_rocks {
		for row_index := range r.round_rocks[column_index] {
			r.rocks_weigth += max_rows - row_index
		}
	}
	return r.rocks_weigth
}
