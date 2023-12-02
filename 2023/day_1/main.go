package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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

func convert_word_to_integer_string(word string) string {
	conversion_map := map[string] int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"zero": 0,
	}
	new_word := word
	if len(word) > 1{
		
		new_value := conversion_map[word]
		new_word = strconv.Itoa(new_value)
	}
	return new_word
}

func parse_answer_two(_data []string){
	number_re := regexp.MustCompile("\\d{1}")
	text_re := []*regexp.Regexp {
		regexp.MustCompile("one|three|four|five|six|seven"),
		regexp.MustCompile("two|seven"),
		regexp.MustCompile("eight|nine"),
	}
	answer_sum := 0
	for _, data_line := range _data {
		number_result := number_re.FindAllStringIndex(data_line, -1)
		text_result := [][]int{}
		for _, regex := range text_re{
			text_result = append(text_result, regex.FindAllStringIndex(data_line, -1)...)
		}
		result := append(number_result, text_result...)
		sort.Slice(result, func(a,b int) bool {
			return result[a][0] < result[b][0]
		})
		
		first := data_line[result[0][0]:result[0][1]]
		last := data_line[result[len(result)-1][0]:result[len(result)-1][1]]
		first = convert_word_to_integer_string(first)
		last = convert_word_to_integer_string(last)
		answer, conversion_error := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		check_error((conversion_error))
		fmt.Println(data_line)
		fmt.Println(answer)
		fmt.Println("------------------")
		answer_sum += answer
	}
	fmt.Println("Answer 2:", answer_sum)
	
}

func main (){
	filepath := os.Args[1]
	data_array := parse_input_from_file(filepath)
	// parse_answer_one(data_array)
	parse_answer_two(data_array)
}


