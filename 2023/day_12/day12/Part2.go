package day12

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_two(_data []string)int{
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := solver(_data, 2)
	fmt.Println("Answer 2:", answer)
	return answer
}