package structs

type Tilemap struct {
	_tiles     [][]*Tile
	Start      *Beam
	StartArray []*Beam
	_max_x     int
	_max_y     int
}

func (tilemap *Tilemap) Init(_max_x, _max_y int) {
	tilemap._tiles = [][]*Tile{}
	for y := 0; y < _max_y; y++ {
		tilemap._tiles = append(tilemap._tiles, []*Tile{})
		for x := 0; x < _max_x; x++ {
			tilemap._tiles[y] = append(tilemap._tiles[y], nil)
		}
	}
	tilemap._max_x = _max_x
	tilemap._max_y = _max_y
}

func (tilemap *Tilemap) Set_tile(_x, _y int, _type string) {
	_position := &Position{X: _x, Y: _y}
	new_tile := &Tile{}
	new_tile.Init(_position, _type)
	tilemap._tiles[_y][_x] = new_tile
	if _x == 0 && _y == 0 {
		tilemap.Start = &Beam{Tile: new_tile, To: RIGHT}
	}
	if _x == 0 {
		tilemap.StartArray = append(tilemap.StartArray, &Beam{Tile: new_tile, To: RIGHT})
	}
	if _y == 0 {
		tilemap.StartArray = append(tilemap.StartArray, &Beam{Tile: new_tile, To: DOWN})
	}
	if _x == tilemap._max_x-1 {
		tilemap.StartArray = append(tilemap.StartArray, &Beam{Tile: new_tile, To: LEFT})
	}
	if _y == tilemap._max_y-1 {
		tilemap.StartArray = append(tilemap.StartArray, &Beam{Tile: new_tile, To: UP})
	}
}

func (tilemap *Tilemap) Link_tiles() {
	for _, values := range tilemap._tiles {
		for _, tile := range values {
			_x := tile.Position.X
			_y := tile.Position.Y
			if _y-1 >= 0 {
				tile._next[UP] = tilemap._tiles[_y-1][_x]
			}
			if _x+1 < tilemap._max_x {
				tile._next[RIGHT] = tilemap._tiles[_y][_x+1]
			}
			if _y+1 < tilemap._max_y {
				tile._next[DOWN] = tilemap._tiles[_y+1][_x]
			}
			if _x-1 >= 0 {
				tile._next[LEFT] = tilemap._tiles[_y][_x-1]
			}
		}
	}
}

func (tilemap *Tilemap) Get_result() int {
	result := 0
	for _, tile_row := range tilemap._tiles {
		for _, tile := range tile_row {
			if tile._energized > 0 {
				result += 1
				tile._energized = 0
				tile._visited_from = []direction{}
			}
		}
	}
	return result
}

func (tilemap *Tilemap) Print_tilemap() {
	for _, tile_row := range tilemap._tiles {
		for _, tile := range tile_row {
			if tile._energized > 0 {
				print("#")
			} else {
				print(tile._type)
			}
		}
		println()
	}
}
