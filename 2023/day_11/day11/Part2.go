package day11

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_two(_data []string)int{
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := game_it(_data, 1_000_000)
	fmt.Println("Answer 2:", answer)
	return answer
}