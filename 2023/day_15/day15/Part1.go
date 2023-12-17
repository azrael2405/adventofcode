package day15

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_one(_data []string) int {
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := parse_data(_data)
	fmt.Println("Answer 1:", answer)
	return answer
}
