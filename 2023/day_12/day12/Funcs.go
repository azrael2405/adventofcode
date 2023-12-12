package day12

import (
	"fmt"
	"helper"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func sum(values []int) int {
	i := 0
	for _, value := range values{
		i += value
	}
	return i
}

func compare_regex_index_with_int(left []int, right int) int{
	switch{
	case left[1] - left[0] < right:
		return -1
	case left[1] - left[0] > right:
		return 1
	default:
		return 0
	}
}


func solve_line(_input_string string, broken_springs []int) int{
	// permutation_count := 0
	permutations := map[string]bool{}
	hash_re := regexp.MustCompile(`#+`)
	for string_index, value := range _input_string{
		value := string(value)
		if string_index == 0{
			if value == "?"{
				permutations["#"] = true
				permutations["."] = true
			} else {
				permutations[value] = true
			}
		}else{
			new_permutations := map[string]bool{}
			next_values := []string{}
			if value == "?"{
				next_values = []string{"#", "."}
			}else{
				next_values = []string{value}
			}
			for permutation := range permutations{
				for _, next_value := range next_values{
					new_permutation := permutation + next_value
					current_springs := hash_re.FindAllStringIndex(new_permutation, -1)
					if len(current_springs) > len(broken_springs){
						continue
					}
					if len(current_springs) < len(broken_springs){
						if len(new_permutation) + sum(broken_springs[len(current_springs):len(broken_springs)-1]) + len(broken_springs[len(current_springs):len(broken_springs)-1])-1 > len(_input_string){
							continue
						}
					}
					if next_value == "."{
						if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) < 0{
							continue
						}
					} else {
						if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) > 0{
							continue
						}
					}
					if string_index == len(_input_string) -1{
						if len(current_springs) != len(broken_springs) {
							continue
						}
						if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) != 0{
							continue
						}
					}

					new_permutations[new_permutation] = true
				}
			}
			// fmt.Println(new_permutations)
			permutations = new_permutations
		}
	}
	
	fmt.Println("----------------")
	fmt.Printf("INPUT: %s\n", _input_string)
	for key := range permutations{
		positions := hash_re.FindAllStringIndex(key, -1)
		fmt.Println(key, "::", positions, "::", broken_springs)
	}
		// for _, permutation := range permutations{
		// 	positions := hash_re.FindAllStringIndex(permutation, -1)
		// 	if len(positions) == len(broken_springs){
			// 		if slices.CompareFunc(positions, broken_springs, ) == 0{
	// 			permutation_count += 1
	// 		}	
	// 	}
	// }
	fmt.Printf("permutations:%d\n", len(permutations))
	return len(permutations)
}


func solver(_data_array []string)int{
	solution := 0
	for _, line := range _data_array{
		split_line := strings.Split(line, " ")
		input_string := split_line[0]
		values_int := []int{}
		for _, value := range strings.Split(split_line[1], ","){
			value_int, conversion_error := strconv.Atoi(value)
			helper.Check_error(conversion_error)
			values_int = append(values_int, value_int)
		}
		solution += solve_line(input_string, values_int)
	}
	return solution
}