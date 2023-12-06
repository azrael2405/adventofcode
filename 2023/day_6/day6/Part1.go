package day6

import (
	"fmt"
	"helper"
	"time"
)




func Parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	races := parse_data(_data)
	answer := parse_winning_moves(races)
	fmt.Println("Answer 1:", answer)
}