package day5

import (
	"helper"
	"strconv"
	"strings"
)

type almanac struct {
	target string
	data []string
}


func parse_data_to_almanac(_data_array []string) map[string]map[int]int{
	seeds := []int{}
	seeds_line := _data_array[0]
	answer_map := map[string]map[int]int{
		"seed": {},
	}
	for _, seed := range strings.Split(strings.TrimSpace(strings.Split(seeds_line, ":")[1]), " "){
		seed_int, conversion_error := strconv.Atoi(seed)
		helper.Check_error(conversion_error)
		seeds = append(seeds, seed_int)
		answer_map["seed"][seed_int] = seed_int
	}
	data_map := map[string]almanac{}
	
	for _, element := range _data_array[1:]{
		element_split := strings.Split(element, "\n")
		name_line := strings.TrimSpace(strings.Split(element_split[0], " ")[0])
		name_array := strings.Split(name_line, "-")
		new_almanac := almanac{
			target: name_array[2],
			data: element_split[1:],
		}
		data_map[name_array[0]] = new_almanac
	}	
	start_category := "seed"
	for current_category:= start_category; current_category != ""; current_category = data_map[current_category].target{
		target_category := data_map[current_category].target
		answer_map[target_category] = map[int]int{}
		for _, data_line := range data_map[current_category].data{
			number_array := strings.Split(data_line, " ")
			destination, conversion_error := strconv.Atoi(number_array[0])
			helper.Check_error(conversion_error)
			source, conversion_error := strconv.Atoi(number_array[1])
			helper.Check_error(conversion_error)
			range_length, conversion_error := strconv.Atoi(number_array[2])
			helper.Check_error(conversion_error)
			for key := range answer_map[current_category]{
				if source <= key && key <= source + range_length - 1{
					offset := key - source
					new_destination := destination + offset
					answer_map[current_category][key] = new_destination
					answer_map[target_category][new_destination] = new_destination
				}
			}
		}
		for _, value := range answer_map[current_category]{
			if answer_map[target_category][value] == 0{
				answer_map[target_category][value] = value
			}
		}
	}
	return answer_map
}

