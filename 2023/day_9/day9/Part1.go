package day9

import (
	"fmt"
	"helper"
	"time"
)


func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	predictions := parse_data_to_values(_data)
	answer := sum(predictions)
	fmt.Println("Answer 1:", answer)
	return answer
}