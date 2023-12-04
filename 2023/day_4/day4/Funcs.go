package day4

import (
	"fmt"
	"helper"
	"os"
	"slices"
	"strings"
)

func parse_data_into_matches_list(_data_array []string) []int{
	points_list := []int{}
	file, error := os.Create("test.out")
	helper.Check_error(error)
	defer file.Close()
	for _, data_line := range _data_array{
		file.WriteString(data_line+"\n")
		card_data  := strings.Split(data_line, ":")
		// card_id, conversion_error := strconv.Atoi(strings.Split(card_data[0], " ")[1])
		card_numbers := strings.Split(card_data[1], "|")
		winning_numbers := []string{}
		number_counter := 0
		for _, number := range strings.Split(strings.TrimSpace(card_numbers[0]), " "){
			number = strings.TrimSpace(number)
			if len(number)>=1{
				winning_numbers = append(winning_numbers, number)
			}
		}
		file.WriteString(strings.Join(winning_numbers, " ")+"\n")
		for _, number := range strings.Split(strings.TrimSpace(card_numbers[1]), " "){
			number = strings.TrimSpace(number)
			if len(number)<1{
				continue
			}
			if slices.Contains(winning_numbers, number){
				file.WriteString(fmt.Sprintf("%s: Winning\n", number))
				number_counter += 1
			}else{
				file.WriteString(fmt.Sprintf("%s: Losing\n", number))
			}
		}
		file.WriteString(fmt.Sprintf("Winning Count: %d\n---------------\n", number_counter))
		points_list = append(points_list, number_counter)
		

	}
	return points_list
}


func pow(x, y int) int{
	if y < 0 {
		panic("y is smaller than 0")
	}
	out := 1
	for i := 0; i < y; i++ {
		out *= x
	}
	return out
}