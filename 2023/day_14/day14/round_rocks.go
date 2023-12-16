package day14

type round_rocks_type struct {
	round_rocks  map[int]map[int]bool
	rocks_weigth int
}

func (r *round_rocks_type) print_map(_max_rows, _max_columns int) {
	for row_index := 0; row_index < _max_rows; row_index++ {
		for column_index := 0; column_index < _max_columns; column_index++ {
			if _, ok := r.round_rocks[column_index][row_index]; ok {
				print("#")
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
	for column_index := range r.round_rocks {
		for row_index := range r.round_rocks[column_index] {
			new_y := s.get_stop_pos_for_round_rock(row_index, column_index)
			r.rocks_weigth += _max_rows - new_y

		}
	}
}
