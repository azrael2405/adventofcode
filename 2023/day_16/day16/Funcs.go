package day16

import "day16/day16/structs"

func parse_input_to_tilemap(_data_array []string) *structs.Tilemap {
	tilemap := &structs.Tilemap{}
	tilemap.Init(len(_data_array), len(_data_array[0]))
	for y, line := range _data_array {
		for x, char := range line {
			tilemap.Set_tile(x, y, string(char))
		}
	}
	tilemap.Link_tiles()
	return tilemap
}

func traverse_tilemap(_tilemap *structs.Tilemap) {
	current_beams := []*structs.Beam{
		{Tile: _tilemap.Start,
			To: structs.RIGHT,
		},
	}
	for {
		if len(current_beams) == 0 {
			break
		}
		current_beam := current_beams[0]
		current_beams = current_beams[1:]
		next_beams := current_beam.Tile.Reflect(current_beam.To)
		current_beams = append(current_beams, next_beams...)
	}
}

func parse_data(_data_array []string) int {
	tilemap := parse_input_to_tilemap(_data_array)
	// println("before traverse\n")
	// tilemap.Print_tilemap()
	traverse_tilemap(tilemap)
	// println("\nafter traverse\n")
	// tilemap.Print_tilemap()
	// println()
	return tilemap.Get_result()
}
