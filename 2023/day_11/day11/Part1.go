package day11

import (
	"fmt"
	"helper"
	"time"
)


func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := game_it(_data, 10)
	fmt.Println("Answer 1:", answer)
	return answer
}