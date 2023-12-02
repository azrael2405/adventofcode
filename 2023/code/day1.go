package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	re := regexp.MustCompile("\\d{1}")
	answer_sum := 0
	for _, data_line := range _data {
		result := re.FindAllStringIndex(data_line, -1)
		first := data_line[result[0][0]:result[0][1]]
		last := data_line[result[len(result)-1][0]:result[len(result)-1][1]]
		answer, conversion_error := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		check_error((conversion_error))
		answer_sum += answer
	}
	fmt.Println("Answer 1:", answer_sum)
}



func main (){

	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	parse_answer_one(data_array)
}
