package day4

import (
	"fmt"
	"helper"
	"time"
)


func Parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := 0
	matches_list := parse_data_into_matches_list(_data)
	for _, matches := range matches_list{
		fmt.Print(matches, " ")
		if matches > 0{
			value := pow(2, matches-1)
			fmt.Print(value)
			answer += value
		}
		fmt.Println("")
	}
	fmt.Println("Answer 1:", answer)
}