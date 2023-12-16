package dayX

import (
	"fmt"
	"helper"
	"time"
)

func Parse_answer_two(_data []string) int {
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := parse_data(_data)
	fmt.Println("Answer 2:", answer)
	return answer
}
