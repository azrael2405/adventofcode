package main

import (
	"fmt"
	"helper"
	"os"
	"strconv"
	"strings"
	"time"
)



func check_error(e error){
	if e != nil{
		panic(e)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	lines := strings.Split(string(file_data), "\n")
	check_error(file_error)
	return lines
}

func game_split(line string) (int, map[string]int){
	line_split := strings.Split(line, ":")
	game_id, conversion_error := strconv.Atoi(strings.TrimSpace(strings.Split(line_split[0], " ")[1]))
	check_error(conversion_error)
	cube_map := map[string]int {
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	for _, game_line := range strings.Split(line_split[1], ";"){
		for _, cube := range strings.Split(game_line, ","){
			cube_split := strings.Split(strings.TrimSpace(cube), " ")
			cube_count, conversion_error := strconv.Atoi(cube_split[0])
			check_error(conversion_error)
			cube_map[cube_split[1]] = Max(cube_map[cube_split[1]], cube_count)
		}
	}
	return game_id, cube_map
}


func parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	game_max := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	answer := 0
	for _, data_line := range _data {
		add_question := true
		game_id, cube_map := game_split(data_line)
		for key, value := range cube_map{
			if value > game_max[key]{
				add_question = false
				break
			}
		}
		if add_question{
			answer += game_id
		}
	}
	fmt.Println("Answer 1:", answer)
}


func parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(),"Answer 2")
	answer := 0
	for _, data_line := range _data {
		_, cube_map := game_split(data_line)
		game_power := 1
		for _, value := range cube_map{
			game_power = game_power * value
		}
		answer += game_power
	}
	fmt.Println("Answer 2:", answer)
	
}

func main (){
	defer helper.TimeTrack(time.Now(), "main")
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)

	parse_answer_one(data_array)
	parse_answer_two(data_array)
}


