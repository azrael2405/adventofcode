package day12

import (
	"fmt"
	"helper"
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

var memo map[string]int = make(map[string] int)
var input_string string = ""
var number_list []int = []int{}
func HandlePanic(key string, r_value int){
	r := recover()
	if r != nil{
		fmt.Println("Panic Recovered:", key, "::", r_value)
		memo[key] = r_value
	} 
}

func solve_recursive(pos_string, pos_number int, r ...int) {
	r_value := 0
	if len(r) > 0{
		r_value = r[0]
	}
	key := fmt.Sprintf("%d_%d", pos_string, pos_number)
	
	if _, ok := memo[key]; ok{
		return
	}
	if pos_string == len(input_string){
		if pos_number == len(number_list){
			memo[key] = 1
			return
		} else {
			memo[key] = 0
			return
		}
	}
	
	if strings.Contains(".?", string(input_string[pos_string])){
		solve_recursive(pos_string+1, pos_number)
		new_key := fmt.Sprintf("%d_%d", pos_string+1, pos_number)
		r_value += memo[new_key]
	}

	if pos_number < len(number_list){
		pos_string_end := pos_string + number_list[pos_number]
		if pos_string_end < len(input_string){
			if strings.Contains("#?", string(input_string[pos_string])) && !strings.Contains(input_string[pos_string:pos_string_end], ".") && string(input_string[pos_string_end]) != "#"{
				solve_recursive(pos_string+1, pos_number+1)
				new_key := fmt.Sprintf("%d_%d", pos_string+1, pos_number+1)
				r_value += memo[new_key]
			}
		}
	}
	memo[key] = r_value
}
	
		

func solve_line(_input_string string, broken_springs []int) int{
	memo = map[string]int{}
	input_string = _input_string
	number_list = broken_springs
	solve_recursive(0, 0)
	fmt.Println(memo)
	fmt.Println("-----")
	output_value := 0
	for _, value := range memo{
		output_value += value
	}
	return output_value
	// permutation_count := 0
	// permutations := map[string]bool{}
	// hash_re := regexp.MustCompile(`#+`)
	// for string_index, value := range _input_string{
	// 	value := string(value)
	// 	if string_index == 0{
	// 		if value == "?"{
	// 			permutations["#"] = true
	// 			permutations["."] = true
	// 		} else {
	// 			permutations[value] = true
	// 		}
	// 	}else{
	// 		new_permutations := map[string]bool{}
	// 		next_values := []string{}
	// 		if value == "?"{
	// 			next_values = []string{"#", "."}
	// 		} else {
	// 			next_values = []string{value}
	// 		}
	// 		for permutation := range permutations{
	// 			for _, next_value := range next_values{
	// 				new_permutation := permutation + next_value
	// 				current_springs := hash_re.FindAllStringIndex(new_permutation, -1)
	// 				if len(current_springs) > len(broken_springs){
	// 					continue
	// 				}
	// 				if len(current_springs) < len(broken_springs){
	// 					if len(new_permutation) + sum(broken_springs[len(current_springs):len(broken_springs)-1]) + len(broken_springs[len(current_springs):len(broken_springs)-1])-1 > len(_input_string){
	// 						continue
	// 					}
	// 				}
	// 				if next_value == "."{
	// 					if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) < 0{
	// 						continue
	// 					}
	// 				} else {
	// 					if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) > 0{
	// 						continue
	// 					}
	// 				}
	// 				if string_index == len(_input_string) -1{
	// 					if len(current_springs) != len(broken_springs) {
	// 						continue
	// 					}
	// 					if slices.CompareFunc(current_springs, broken_springs[:len(current_springs)], compare_regex_index_with_int) != 0{
	// 						continue
	// 					}
	// 				}
	// 				new_permutations[new_permutation] = true
	// 			}
	// 		}
	// 		permutations = new_permutations
	// 	}
	// }
	// return len(permutations)
}


func solver(_data_array []string, part int)int{
	solution := 0
	for _, line := range _data_array{
		split_line := strings.Split(line, " ")
		input_string := split_line[0]
		value_string := split_line[1]
		values_int := []int{}
		if part == 2{
			value_string_array := []string{}
			input_string_array := []string{}
			for i:= 0; i < 5; i++{
				value_string_array = append(value_string_array, value_string)
				input_string_array = append(input_string_array, input_string)
			}
			value_string = strings.Join(value_string_array, ",")
			input_string = strings.Join(input_string_array, "?")
		}
		for _, value := range strings.Split(value_string, ","){
			value_int, conversion_error := strconv.Atoi(value)
			helper.Check_error(conversion_error)
			values_int = append(values_int, value_int)
		}
		solution += solve_line(input_string, values_int)
	}
	return solution
}