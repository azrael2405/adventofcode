package day7

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_two(_data []string) int {
	defer helper.TimeTrack(time.Now(), "Answer 2")
	hands := parse_input_to_hands(_data, 2)
	hands = sort_hands(hands,2)
	answer := count_answer(hands)
	fmt.Println("Answer 2:", answer)
	return answer
}