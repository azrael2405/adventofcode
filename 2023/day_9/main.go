package main

import (
	"day9/day9"
	"helper"
	"os"
	"strings"
	"time"
)

func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	helper.Check_error(file_error)
	lines := strings.Split(string(file_data), "\n")
	return lines
}


func main (){
	defer helper.TimeTrack(time.Now(), "main")
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	day9.Parse_answer_one(data_array)
	day9.Parse_answer_two(data_array)
}


