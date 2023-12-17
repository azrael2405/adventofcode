package day15

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_two(_data []string) int {
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := parse_data_part_two(_data)
	fmt.Println("Answer 2:", answer)
	return answer
}
