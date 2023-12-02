package main

import (
	"fmt"
	"helper"
	"os"
	"strings"
	"time"
)

func check_error(e error){
	if e != nil{
		panic(e)
	}
}

func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	lines := strings.Fields(string(file_data))
	check_error(file_error)
	return lines
}


func parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := ""
	fmt.Println("Answer 1:", answer)
}


func parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := ""
	fmt.Println("Answer 2:", answer)
	
}

func main (){
	defer helper.TimeTrack(time.Now(), "main")
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	parse_answer_one(data_array)
	parse_answer_two(data_array)
}


