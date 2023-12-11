package day11

import (
	"fmt"
	"regexp"
)


type position_type struct {
	key int
	x int
	y int
	x_expansion int
	y_expansion int
}

func (p position_type) to_string() string{
	return fmt.Sprintf("%2d: %5d + %2d / %5d + %2d", p.key, p.x, p.x_expansion, p.y, p.y_expansion)
}

func abs(value int)int{
	if value < 0{
		value *= -1
	}
	return value
}

func sum(values []int) int{
	sum_value := 0
	for _, value := range values{
		sum_value += value
	}
	return sum_value
}


func find_galaxies(_data_array []string) []*position_type{
	galaxies := []*position_type{}
	empty_re := regexp.MustCompile(`#{1}`)
	key := 1
	for x_index, data_line := range _data_array{
		found_string := empty_re.FindAllStringIndex(data_line, -1)
		for _, entry := range found_string{
			galaxies = append(galaxies, &position_type{
				key: key,
				x: x_index,
				y: entry[0],
				x_expansion: 0,
				y_expansion: 0,
			})
			key++
		}
	}
	return galaxies
}


func find_empty_spaces(_galaxies []*position_type, max_rows, max_cols int) map[string]map[int]bool{
	empty_rows := make(map[int]bool, max_rows)
	empty_cols := make(map[int]bool, max_rows)
	for x := 0; x < max_rows; x++{
		empty_rows[x] = true
	}
	for y := 0; y < max_cols; y++{
		empty_cols[y] = true
	}
	for _, galaxy := range _galaxies{
		delete(empty_rows, galaxy.x)
		delete(empty_cols, galaxy.y)
	}

	return map[string]map[int]bool{
		"empty_rows": empty_rows,
		"empty_cols": empty_cols,
	}
}

func expand_galaxies(_galaxies []*position_type, empty_rows, empty_cols map[int]bool)[]*position_type{
	for index, galaxy := range _galaxies{
		count_rows := 0
		count_cols := 0
		if len(empty_rows) > 0{
			for key := range empty_rows{
				if key < galaxy.x{
					count_rows += 1
				}
			}
		}
		if len(empty_cols) > 0{
			for key := range empty_cols{
				if key < galaxy.y{
					count_cols += 1
				}
			}
		}

		galaxy.x_expansion = count_rows
		galaxy.y_expansion = count_cols
		_galaxies[index] = galaxy
	}
	return _galaxies
}

func find_shortest_routes_between_galaxies(_galaxies []*position_type, expansion int)[]int{
	shortest_routes := []int{}
	found_pairs := map[*position_type]bool{}
	for _, galaxy := range _galaxies{
		for _, next_galaxy := range _galaxies{
			if found_pairs[next_galaxy] || next_galaxy == galaxy {
				continue
			}
			new_path := abs(galaxy.x-next_galaxy.x) + abs(galaxy.y-next_galaxy.y) + expansion * abs(galaxy.x_expansion - next_galaxy.x_expansion) + expansion * abs(galaxy.y_expansion - next_galaxy.y_expansion)
			shortest_routes = append(shortest_routes, new_path)
			// fmt.Println(galaxy.key, next_galaxy.key, "::", new_path)
		}
		found_pairs[galaxy] = true
	}
	return shortest_routes
}

func print_galaxies(_galaxies []*position_type){
	for _, galaxy := range _galaxies{
		fmt.Println(galaxy.to_string())
	}
}

func game_it(_data_array []string, expansion int) int{
	galaxies := find_galaxies(_data_array)
	empty_spaces := find_empty_spaces(galaxies, len(_data_array), len(_data_array[0]))
	galaxies = expand_galaxies(galaxies, empty_spaces["empty_rows"], empty_spaces["empty_cols"])
	print_galaxies(galaxies)
	shortest_paths := find_shortest_routes_between_galaxies(galaxies, expansion)
	
	return sum(shortest_paths)
}