package day5

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

type almanac struct {
	target string
	data []string
}




func parse_data_to_almanac(_data_array []string, seeds []int) map[string]map[int]int{
	answer_map := map[string]map[int]int{
		"seed": {},
	}
	for _, seed := range seeds{
		answer_map["seed"][seed] = seed
	}
	data_map := map[string]almanac{}
	
	for _, element := range _data_array{
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


func parse_data_to_almanac_with_seed_ranges(_data_array []string, seeds [][]int) map[string]map[string][]int{
	answer_map := map[string]map[string][]int{
		"seed": {},
	}
	for _, seed_range := range seeds{

		answer_map["seed"][fmt.Sprintf("%d-%d", seed_range[0], seed_range[1])] = seed_range
	}
	data_map := map[string]almanac{}
	
	for _, element := range _data_array{
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
		answer_map[target_category] = map[string][]int{}
		for _, data_line := range data_map[current_category].data{
			number_array := strings.Split(data_line, " ")
			destination, conversion_error := strconv.Atoi(number_array[0])
			helper.Check_error(conversion_error)
			source, conversion_error := strconv.Atoi(number_array[1])
			helper.Check_error(conversion_error)
			range_length, conversion_error := strconv.Atoi(number_array[2])
			helper.Check_error(conversion_error)
			for key, old_range := range answer_map[current_category]{
				range_start, conversion_error := strconv.Atoi(strings.Split(key, "-")[0])
				helper.Check_error(conversion_error)
				range_end, conversion_error := strconv.Atoi(strings.Split(key, "-")[1])
				helper.Check_error(conversion_error)
				source_end := source + range_length - 1
				destination_end := destination + range_length -1
				switch {
				case source <= range_start && range_end <= source_end:
					offset_start := range_start - source
					offset_end := range_end - source
					answer_map[current_category][key] = []int{destination+offset_start, destination+offset_end}
				case source <= range_start && range_start <= source_end && range_end > source_end:
					delete(answer_map[current_category], key)
					new_key1 := fmt.Sprintf("%d-%d", range_start, source_end)
					new_key2 := fmt.Sprintf("%d-%d", source_end+1, range_end)
					offset := range_start - source
					new_data1 := []int{destination+offset, destination_end}
					new_data2 := []int{old_range[1] + range_end-source_end+1, old_range[1]}
					answer_map[current_category][new_key1] = new_data1
					answer_map[current_category][new_key2] = new_data2
					if new_data1[0] <0 ||new_data1[1] <0 || new_data2[0] < 0|| new_data2[1] < 0{
						fmt.Println(new_data1,";;", new_data2)
						panic("Wrong numbers!")
					}
				case source <= range_end && range_end <= source_end && range_start < source:
					delete(answer_map[current_category], key)
					new_key1 := fmt.Sprintf("%d-%d", source, range_end)
					new_key2 := fmt.Sprintf("%d-%d", range_start, source-1)
					offset := range_end - source
					new_data1 := []int{destination, destination+offset}
					new_data2 := []int{old_range[0], old_range[0] + (source - 1 - range_start)}
					answer_map[current_category][new_key1] = new_data1
					answer_map[current_category][new_key2] = new_data2
					if new_data1[0] <0 ||new_data1[1] <0 || new_data2[0] < 0|| new_data2[1] < 0{
						fmt.Println(new_data1,";;", new_data2)
						panic("Wrong numbers!")
					}
				case range_start < source && source_end < range_end:
					delete(answer_map[current_category], key)
					new_key1 := fmt.Sprintf("%d-%d", source, source_end)
					new_key2 := fmt.Sprintf("%d-%d", range_start, source-1)
					new_key3 := fmt.Sprintf("%d-%d", source_end + 1, range_end)
					new_data1 := []int{destination, destination+range_length-1}
					new_data2 := []int{old_range[0], old_range[0] + (source - 1 - range_start)}
					new_data3 := []int{old_range[1] - source_end + range_end + 1, old_range[1]}
					answer_map[current_category][new_key1] = new_data1
					answer_map[current_category][new_key2] = new_data2
					answer_map[current_category][new_key3] = new_data3
					if new_data1[0] <0 ||new_data1[1] <0 || new_data2[0] < 0|| new_data2[1] < 0|| new_data3[0] < 0|| new_data3[1] < 0{
						fmt.Println(new_data1,";;", new_data2, ";;", new_data3)
						panic("Wrong numbers!")
					}
				}
			}
		}
		for _, value := range answer_map[current_category]{
			new_key := fmt.Sprintf("%d-%d", value[0], value[1])
			answer_map[target_category][new_key] = value
		}
	}
	return answer_map
}

