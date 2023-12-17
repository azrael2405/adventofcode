package day16

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_one(_data []string) int {
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer := parse_data(_data, 1)
	fmt.Println("Answer 1:", answer)
	return answer
}
