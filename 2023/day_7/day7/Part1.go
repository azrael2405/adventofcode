package day7

import (
	"fmt"
	"helper"
	"time"
)



func Parse_answer_one(_data []string) int{
	defer helper.TimeTrack(time.Now(), "Answer 1")
	hands := parse_input_to_hands(_data, 1)
	hands = sort_hands(hands, 1)
	answer := count_answer(hands)
	fmt.Println("Answer 1:", answer)
	return answer
}