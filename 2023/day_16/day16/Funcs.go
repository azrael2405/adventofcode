package day16

import (
	"day16/day16/structs"
)

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

func traverse_tilemap_part1(_tilemap *structs.Tilemap) {
	current_beams := []*structs.Beam{
		_tilemap.Start,
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

func traverse_tilemap_part2(_tilemap *structs.Tilemap) int {
	current_highest := 0
	var current_highest_Beam *structs.Beam
	for _, beam := range _tilemap.StartArray {
		current_beams := []*structs.Beam{
			beam,
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
		result := _tilemap.Get_result()
		// println("current_beam", beam.Tile.Position.X, "-", beam.Tile.Position.Y, "::", beam.To, "::", result)
		if result > current_highest {
			current_highest = result
			current_highest_Beam = beam
		}
	}
	println("current_highest_Beam", current_highest_Beam.Tile.Position.X, "-", current_highest_Beam.Tile.Position.Y, "::", current_highest_Beam.To)
	return current_highest
}

func parse_data(_data_array []string, part int) int {
	tilemap := parse_input_to_tilemap(_data_array)
	if part == 1 {
		traverse_tilemap_part1(tilemap)
		return tilemap.Get_result()
	} else {
		return traverse_tilemap_part2(tilemap)

	}
}
