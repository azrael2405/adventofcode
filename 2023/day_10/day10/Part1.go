package day10

import (
	"fmt"
	"helper"
	"time"
)


func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := run_game(_data)
	fmt.Println("Answer 1:", answer)
	return answer
}