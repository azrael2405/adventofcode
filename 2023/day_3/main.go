package main

import (
	"day3/day3"
	"helper"
	"os"
	"strings"
	"time"
)


func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	day3.Check_error(file_error)
	lines := strings.Fields(string(file_data))
	return lines
}


func main (){
	defer helper.TimeTrack(time.Now(), "main")
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	x_max := len(data_array)
	y_max := len(data_array[0])
	day3.Set_my_globals(x_max, y_max)
	day3.Parse_answer_one(data_array)
	day3.Parse_answer_two(data_array)
}
