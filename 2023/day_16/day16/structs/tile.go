package structs

import "slices"

type Tile struct {
	_energized    int
	Position      *Position
	_type         string
	_next         map[direction]*Tile
	_visited_from []direction
}

func (tile *Tile) Init(_position *Position, _type string) {
	tile._energized = 0
	tile.Position = _position
	tile._type = _type
	tile._next = make(map[direction]*Tile)
	tile._visited_from = make([]direction, 0)
}

func (tile *Tile) get_next(_direction direction) *Beam {
	if _, ok := tile._next[_direction]; ok {
		// println("get_next", tile._position.X, "-", tile._position.Y, "::", _direction, "::", tile._type, "::", tile._next[_direction]._position.X, "-", tile._next[_direction]._position.Y)
		return &Beam{Tile: tile._next[_direction], To: _direction}
	} else {
		// println("get_next", tile._position.X, "-", tile._position.Y, "::", _direction, "::", "nil")
		return nil
	}
}

func (tile *Tile) Reflect(_to direction) []*Beam {
	next_beams := []*Beam{}
	already_visited := slices.Contains(tile._visited_from, _to)
	if already_visited {
		return []*Beam{}
	}
	tile._visited_from = append(tile._visited_from, _to)
	tile._energized += 1
	switch tile._type {
	case ".":
		next_tile := tile.get_next(_to)
		if next_tile != nil {
			next_beams = append(next_beams, next_tile)
		}
	case "/":
		switch _to {
		case UP:
			next_tile := tile.get_next(RIGHT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case RIGHT:
			next_tile := tile.get_next(UP)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case DOWN:
			next_tile := tile.get_next(LEFT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case LEFT:
			next_tile := tile.get_next(DOWN)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		}
	case "\\":
		switch _to {
		case UP:
			next_tile := tile.get_next(LEFT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case RIGHT:
			next_tile := tile.get_next(DOWN)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case DOWN:
			next_tile := tile.get_next(RIGHT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case LEFT:
			next_tile := tile.get_next(UP)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		}
	case "|":
		switch _to {
		case RIGHT:
			next_tile := tile.get_next(UP)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
			next_tile = tile.get_next(DOWN)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case LEFT:
			next_tile := tile.get_next(UP)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
			next_tile = tile.get_next(DOWN)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case UP:
			next_tile := tile.get_next(_to)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case DOWN:
			next_tile := tile.get_next(_to)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		}
	case "-":
		switch _to {
		case UP:
			next_tile := tile.get_next(LEFT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
			next_tile = tile.get_next(RIGHT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case DOWN:
			next_tile := tile.get_next(LEFT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
			next_tile = tile.get_next(RIGHT)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case RIGHT:
			next_tile := tile.get_next(_to)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		case LEFT:
			next_tile := tile.get_next(_to)
			if next_tile != nil {
				next_beams = append(next_beams, next_tile)
			}
		}
	}
	return next_beams
}
