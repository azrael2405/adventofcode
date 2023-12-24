package structs

import (
	"helper"
	"slices"
)

type HoleMap struct {
	Holes   [][]*Hole
	Top_row int
}

func (hm *HoleMap) Init() {
	hm.Holes = [][]*Hole{{}}
	hm.Top_row = 0
	hm.Holes = append(hm.Holes, []*Hole{{}})
}

func (hm *HoleMap) DigTrench(_start_position *Position, _direction Direction, _length int, color string) *Position {
	println(_start_position.X, _start_position.Y, _direction, _length, color)
	end_position := &Position{_start_position.X, _start_position.Y}
	switch _direction {
	case UP:
		end_position.Y -= _length
		if end_position.Y < hm.Top_row {
			hm.add_rows(helper.Abs(hm.Top_row-end_position.Y), UP)
		}
		adjusted_y := end_position.Y - hm.Top_row
		for i := 0; i < _length; i++ {
			new_hole := &Hole{
				Start:  &Position{_start_position.X, end_position.Y + i},
				End:    &Position{_start_position.X, end_position.Y + i},
				Length: 1,
				Color:  color,
			}
			hm.Holes[adjusted_y+i] = append(hm.Holes[adjusted_y+i], new_hole)
		}
	case RIGHT:
		end_position.X += _length
		adjusted_y := end_position.Y - hm.Top_row
		new_hole := &Hole{Start: _start_position, End: end_position, Length: _length, Color: color}
		hm.Holes[adjusted_y] = append(hm.Holes[adjusted_y], new_hole)
	case DOWN:
		end_position.Y += _length
		adjusted_y := end_position.Y - hm.Top_row
		println(end_position.Y, hm.Top_row, len(hm.Holes), adjusted_y-len(hm.Holes)+1)
		if end_position.Y >= len(hm.Holes) {
			hm.add_rows(adjusted_y-len(hm.Holes)+1, DOWN)
		}
		for i := 0; i < _length; i++ {
			new_hole := &Hole{
				Start:  &Position{_start_position.X, _start_position.Y + i},
				End:    &Position{_start_position.X, _start_position.Y + i},
				Length: 1,
				Color:  color,
			}
			hm.Holes[adjusted_y-i] = append(hm.Holes[adjusted_y-i], new_hole)
		}
	case LEFT:
		end_position.X -= _length
		adjusted_y := end_position.Y - hm.Top_row
		new_hole := &Hole{Start: end_position, End: _start_position, Length: _length, Color: color}
		hm.Holes[adjusted_y] = append(hm.Holes[adjusted_y], new_hole)
	}
	return end_position
}

func (hm *HoleMap) add_rows(_num_rows int, _direction Direction) {
	if _direction == UP {
		for i := 0; i < _num_rows; i++ {
			hm.Holes = append([][]*Hole{{}}, hm.Holes...)
		}
		hm.Top_row -= _num_rows
	} else if _direction == DOWN {
		for i := 0; i < _num_rows; i++ {
			hm.Holes = append(hm.Holes, []*Hole{{}})
		}
	}
}

func (hm *HoleMap) DigHole() int {
	answer := 0

	for _, row := range hm.Holes {
		println("row: ", len(row))
		for _, hole := range row {
			println(hole.String())
		}
		if len(row) == 0 {
			continue
		}
		slices.SortFunc(
			row,
			func(_hole1 *Hole, _hole2 *Hole) int {
				if _hole1.Start.X < _hole2.Start.X {
					return -1
				} else if _hole1.Start.X > _hole2.Start.X {
					return 1
				}
				return 0

			},
		)
		var start_hole *Hole = nil
		var end_hole *Hole = nil
		for _, hole := range row {
			if start_hole == nil {
				start_hole = hole
				answer += hole.Length
				continue
			}
			if end_hole == nil {
				end_hole = hole
				answer += hole.Length
				answer += end_hole.Start.X - start_hole.End.X
				start_hole = nil
				end_hole = nil
			}
		}
	}
	return answer
}
