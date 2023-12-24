package day18

import (
	"day18/day18/structs"
	"helper"
	"strconv"
	"strings"
)

func dig_trenches(_data_array []string) *structs.HoleMap {
	hole_map := &structs.HoleMap{}
	hole_map.Init()
	start_point := &structs.Position{X: 0, Y: 0}
	for _, line := range _data_array {
		line_split := strings.Split(line, " ")
		direction := structs.Direction_from_char(line_split[0])
		length, ok := strconv.Atoi(line_split[1])
		helper.Check_error(ok)
		color := line_split[2]
		start_point = hole_map.DigTrench(start_point, direction, length, color)
	}
	return hole_map
}

func parse_data(_data_array []string) int {
	hole_map := dig_trenches(_data_array)
	answer_value := hole_map.DigHole()
	return answer_value
}
