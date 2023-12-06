package day6

import (
	"fmt"
	"helper"
	"strings"
	"time"
)

func Parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 2")
	fixed_data := []string{
		strings.Replace(_data[0], " ", "", -1),
		strings.Replace(_data[1], " ", "", -1),
	}
	races := parse_data(fixed_data)
	answer := parse_winning_moves(races)
	fmt.Println("Answer 2:", answer)
}